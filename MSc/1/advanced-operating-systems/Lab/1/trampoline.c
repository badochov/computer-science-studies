#include <stdio.h>
#include <sys/mman.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <stdint.h>
#include <malloc.h>

typedef void (*formatter)(int);

uint8_t printf_template[] = {
    0x48, 0xba, // movabs rdx, printf
};

uint8_t char_template[] = {
    0x89, 0xfe, // mov esi, edi
    0x31, 0xc0, // xor eax, eax
    0x48, 0xbf  // movabs rdi, ptr
};

int (*p)(const char *__restrict__ __format, ...) = printf;

uint8_t post_template[] = {
    0xff, 0xe2 // jmp rdx
};

formatter make_formatter(const char *format)
{
    size_t s = sizeof(printf_template) + sizeof(&p) + sizeof(char_template) + sizeof(&format) + sizeof(post_template);
    uint8_t *data = valloc(s);
    if (data == NULL)
    {
        return NULL;
    }
    (void)memcpy(data, printf_template, sizeof(printf_template));
    (void)memcpy(data + sizeof(printf_template), &p, sizeof(p));
    (void)memcpy(data + sizeof(printf_template) + sizeof(&p), char_template, sizeof(char_template));
    (void)memcpy(data + sizeof(printf_template) + sizeof(&p) + sizeof(char_template), &format, sizeof(format));
    (void)memcpy(data + sizeof(printf_template) + sizeof(&p) + sizeof(char_template) + sizeof(&format), post_template, sizeof(post_template));

    if (mprotect(data, s, PROT_EXEC | PROT_READ | PROT_WRITE) != 0)
    {
        return NULL;
    }

    return (formatter)data;
}

int main()
{
    formatter x08_format = make_formatter("%08x\n");
    formatter xalt_format = make_formatter("%#x\n");
    formatter d_format = make_formatter("%d\n");
    formatter verbose_format = make_formatter("Liczba: %9d!\n");

    x08_format(0x1234);
    xalt_format(0x5678);
    d_format(0x9abc);
    verbose_format(0xdef0);

    return 0;
}