#include "accelerometer.h"
#include "i2c_init.h"
#include "leds.h"
#include "watchdog.h"
#include <gpio.h>
#include <irq.h>
#include <stddef.h>
#include <stdlib.h>
#include <stm32.h>

#define CTRL_REG1 0x20
#define CTRL_REG3 0x22

#define OUT_X 0x29
#define OUT_Y 0x2B
#define OUT_Z 0x2D

#define SCL_GPIO GPIOB
#define SCL_LINE 8
#define SDA_GPIO GPIOB
#define SDA_LINE 9

#define PCLK1_HZ 16000000

#define X_LED RED
#define Y_LED GREEN
#define Z_LED BLUE

#define LIS35DE_ADDR 0x1C
// Power on all axes enabled.
#define CTRL_REG1_INIT_VALUE 0b01000111
// Push pull , interrupt level high, interrupt on data ready on INT1.
#define CTRL_REG3_INIT_VALUE 0b00000100

#define INT1_GPIO GPIOA
#define INT1_LINE 1
#define INT1_MASK 1u << INT1_LINE

// Typeof I2C register.
typedef int8_t reg_t;
// Typeof I2C register value.
typedef int8_t reg_val_t;

// Typeof function that will be called when write to given register finished.
typedef void (*i2c_write_callback)();
// Typeof function that will be called when read from given register finished.
// Val is read value.
typedef void (*i2c_read_callback)(reg_val_t val);

// Struct to save I2C state.
typedef struct {
    enum i2c_mode_enum { READ, WRITE } mode;
    reg_t reg;
    reg_val_t value;
    union {
        i2c_write_callback write;
        i2c_read_callback read;
    } callback;
} i2c_data_t;

i2c_data_t i2c_data;

static void init_leds() {
    leds_setup();
    Green2LEDon();
}

static void configure_gpio() {
    GPIOafConfigure(SCL_GPIO, SCL_LINE, GPIO_OType_OD, GPIO_Low_Speed, GPIO_PuPd_NOPULL,
                    GPIO_AF_I2C1);
    GPIOafConfigure(SDA_GPIO, SDA_LINE, GPIO_OType_OD, GPIO_Low_Speed, GPIO_PuPd_NOPULL,
                    GPIO_AF_I2C1);
}

// Writes value to given register of accelerometer. Callbacks callback on success.
static void i2c_write(reg_t reg, reg_val_t value, i2c_write_callback callback) {
    i2c_data.mode = WRITE;
    i2c_data.reg = reg;
    i2c_data.value = value;
    i2c_data.callback.write = callback;

    I2C1->CR2 |= I2C_CR2_ITBUFEN;
    I2C1->CR1 |= I2C_CR1_START;
}

// Reads from given register of accelerometer. Callbacks callback on success with read value.
static void i2c_read(reg_t reg, i2c_read_callback callback) {
    i2c_data.mode = READ;
    i2c_data.reg = reg;
    i2c_data.callback.read = callback;

    I2C1->CR2 |= I2C_CR2_ITBUFEN;
    I2C1->CR1 |= I2C_CR1_START;
}

// Converts value read from register to LED intensity.
static int val_to_intensity(reg_val_t val) {
    return abs(val);
}

// Function called whenever new accelerometer measurement is ready.
static void handle_data_ready();

void EXTI1_IRQHandler(void) {
    Green2LEDoff();
    if (EXTI->PR & EXTI_PR_PR1) {
        EXTI->PR = EXTI_PR_PR1;
        handle_data_ready();
    }
    Green2LEDon();
}

// Handles value of measurement of Z axis, which is last.
// Checks if there is new data to handle.
static void handle_z(reg_val_t z) {
    leds_change_blue_intensity(val_to_intensity(z));
    if ((INT1_GPIO->IDR & INT1_MASK) && !(EXTI->PR & EXTI_PR_PR1)) {
        handle_data_ready();
    }
}

// Reads measurement of z axis.
// Gets measurement of y axis and changes corresponding LED's intensity.
static void read_z(reg_val_t y) {
    leds_change_green_intensity(val_to_intensity(y));

    i2c_read(OUT_Z, &handle_z);
}

// Reads measurement of y axis.
// Gets measurement of x axis and changes corresponding LED's intensity.
static void read_y(reg_val_t x) {
    leds_change_red_intensity(val_to_intensity(x));
    i2c_read(OUT_Y, &read_z);
}

static void handle_data_ready() {
    i2c_read(OUT_X, &read_y);
}

// Sets REG1 of accelerometer to start measurement.
static void set_reg1_config() {
    i2c_write(CTRL_REG1, CTRL_REG1_INIT_VALUE, NULL);
}

// Sets REG3 of accelerometer to setup interruptions. Enables interruptions.
static void set_reg3_config() {
    GPIOinConfigure(INT1_GPIO, INT1_LINE, GPIO_PuPd_DOWN, EXTI_Mode_Interrupt, EXTI_Trigger_Rising);
    NVIC_EnableIRQ(EXTI1_IRQn);
    EXTI->PR = EXTI_PR_PR1;

    i2c_write(CTRL_REG3, CTRL_REG3_INIT_VALUE, &set_reg1_config);
}

static void clear_reg3() {
    i2c_write(CTRL_REG3, 0, &set_reg3_config);
}

// Action to be performed whenever fatal error is encountered.
static void panic() {
    NVIC_SystemReset();
}

// Handler for I2C read.
void I2C1_EV_IRQHandler_read() {
    enum next { SB, ADDR, BTF, SB2, ADDR2, RXNE };
    static enum next next_state = SB;

    uint16_t sr1 = I2C1->SR1;
    if (next_state == SB && (sr1 & I2C_SR1_SB)) {
        I2C1->DR = LIS35DE_ADDR << 1;
    } else if (next_state == ADDR && (sr1 & I2C_SR1_ADDR)) {
        I2C1->SR2;
        I2C1->DR = i2c_data.reg;
    } else if (next_state == BTF && (sr1 & I2C_SR1_BTF)) {
        I2C1->CR1 |= I2C_CR1_START;
    } else if (next_state == SB2 && (sr1 & I2C_SR1_SB)) {
        I2C1->DR = LIS35DE_ADDR << 1 | 1;

        I2C1->CR1 &= ~I2C_CR1_ACK;
    } else if (next_state == ADDR2 && (sr1 & I2C_SR1_ADDR)) {
        I2C1->SR2;

        I2C1->CR1 |= I2C_CR1_STOP;
    } else if (next_state == RXNE && (sr1 & I2C_SR1_RXNE)) {
        I2C1->CR2 ^= I2C_CR2_ITBUFEN;
        if (i2c_data.callback.read != NULL) {
            i2c_data.callback.read(I2C1->DR);
        }
    } else {
        return;
    }
    watchdog_reload();
    next_state++;
    if (next_state > RXNE) {
        next_state = 0;
    }
}

// Handler for I2C write.
void I2C1_EV_IRQHandler_write() {
    enum next { SB, ADDR, TX, BTF };
    static enum next next_state = SB;

    uint16_t sr1 = I2C1->SR1;
    if (next_state == SB && (sr1 & I2C_SR1_SB)) {
        I2C1->DR = LIS35DE_ADDR << 1;
    } else if (next_state == ADDR && (sr1 & I2C_SR1_ADDR)) {
        I2C1->SR2;
        I2C1->DR = i2c_data.reg;
    } else if (next_state == TX && (sr1 & I2C_SR1_TXE)) {
        I2C1->DR = i2c_data.value;
    } else if (next_state == BTF && (sr1 & I2C_SR1_BTF)) {
        I2C1->CR1 |= I2C_CR1_STOP;
        I2C1->CR2 ^= I2C_CR2_ITBUFEN;
        if (i2c_data.callback.write != NULL) {
            i2c_data.callback.write();
        }
    } else {
        return;
    }
    watchdog_reload();
    next_state++;
    if (next_state > BTF) {
        next_state = 0;
    }
}

void I2C1_EV_IRQHandler() {
    Green2LEDoff();
    if (i2c_data.mode == READ) {
        I2C1_EV_IRQHandler_read();
    } else {
        I2C1_EV_IRQHandler_write();
    }
    Green2LEDon();
}

void I2C1_ER_IRQHandler() {
    Green2LEDoff();
    panic();
}

void accelerometer_setup() {
    init_leds();

    i2c_init(PCLK1_HZ);
    configure_gpio();
}

void accelerometer_start() {
    watchdog_start(1000);
    i2c_write(CTRL_REG1, 0, &clear_reg3);
}