SYS_EXIT EQU 60
SYS_READ EQU 0
SYS_WRITE EQU 1

STR_END EQU 0

ZERO_CHAR EQU '0'
NINE_CHAR EQU '9'

MOD_VAL EQU 0x10FF80
UNICODE_START EQU 0x80

IN_BUF_SIZE EQU 4096
OUT_BUF_SIZE EQU 4096

UTF8_1BYTE EQU 0b10000000
UTF8_2BYTES EQU 0b11000000
UTF8_3BYTES EQU 0b11100000
UTF8_4BYTES EQU 0b11110000
UTF8_MAX EQU 0b11110111

UTF8_MIN_VAL_2BYTES EQU 0x80
UTF8_MIN_VAL_3BYTES EQU 0x800
UTF8_MIN_VAL_4BYTES EQU 0x10000

UTF8_MAX_UNICODE_VALUE EQU 0x10FFFF

UTF8_BYTE_MASK EQU 0b10000000
UTF8_ADDITIONAL_BYTE_MAX EQU 0b10111111

LOWER6_BITS EQU 0b111111

CHAR_MASK EQU 0xff

STDIN EQU 0
STDOUT EQU 1

EOF_CHAR EQU 0xff
EOF EQU -1

GLOBAL _exit
GLOBAL parse_args
GLOBAL calculate_diacritization_polynomial
GLOBAL str_to_unsigned
GLOBAL read_char_buffered
GLOBAL write_char_buffered
GLOBAL flush_out_buff
GLOBAL read_additional_bytes
GLOBAL read_utf8
GLOBAL print_utf8
GLOBAL _start

section .bss
    in_buf:
            resb IN_BUF_SIZE
    out_buf:
            resb OUT_BUF_SIZE
    in_buf_reads_cnt:
            resd 1                         ; counter of how many arguments are in the in buffer
    in_buf_reads_size:
            resd 1                         ; counter of how many arguments are in the in buffer
    out_buf_cnt:
            resd 1                         ; counter of how many arguments are in the out buffer
    argc:
            resd 1                         ; argc
    coefficients_ptr:
            resq 1                         ; pointer to polynomial coefficients

section .text

; entry point
_start:
        mov edi, [rsp]
        lea rsi, [rsp + 8]

        mov [argc], edi
        dec DWORD [argc]

        call parse_args

        mov [coefficients_ptr], rax

; loop reading characters till EOF is encountered
read_loop:
        call read_utf8
        ; check if we got EOF and loop should be ended
        cmp eax, EOF
        je read_loop_end
        ; check if we got standard ascii character
        cmp eax, UNICODE_START
        jl print_calculated
        ; handle non ascii character
        sub eax, UNICODE_START
        ; calculate polynomial
        mov edi, eax
        mov rsi, [coefficients_ptr]
        mov edx, [argc]

        call calculate_diacritization_polynomial

        add eax, UNICODE_START
; prints unicode entity
print_calculated:

        mov edi, eax

        call print_utf8

        jmp read_loop
; end of the read loop
read_loop_end:
        ; return from the program with code 0
        xor edi, edi
        jmp _exit

; prints unicode character as utf-8 entity
; exits the program with code 1 on error
; edi - unicode character
print_utf8:
        push rbx                       ; to store number of additional bytes to write
        push r12                       ; to store original character

        mov r12d, edi

        ; check if unicode should be encoded on 4 bytes
        cmp edi, UTF8_MIN_VAL_4BYTES
        jl pu8_if_3b

; set first byte for write
        shr edi, 18
        or edi, UTF8_4BYTES

        mov ebx, 12 ; to signal 3 more bytes need to be written

        jmp pu8_fi
; check if unicode should be encoded on 3 bytes
pu8_if_3b:
        cmp edi, UTF8_MIN_VAL_3BYTES
        jl pu8_if_2b

; set first byte for write
        shr edi, 12
        or edi, UTF8_3BYTES

        mov ebx, 6 ; to signal 2 more bytes need to be written

        jmp pu8_fi
; check if unicode should be encoded on 2 bytes
pu8_if_2b:
        cmp edi, UTF8_MIN_VAL_2BYTES
        jl pu_fn_end

; set first byte for write
        shr edi, 6
        or edi, UTF8_2BYTES

        mov ebx, 0 ; to signal 1 more byte needs to be written

pu8_fi:
        call write_char_buffered
; prepare next byte
        mov edi, r12d
        mov cl, bl
        shr edi, cl                    ; get current 6 bits
        and edi, LOWER6_BITS           ; limit to only lower 6 bits
        or edi, UTF8_BYTE_MASK         ; add mask

        sub ebx, 6                     ; move to next 6 bytes
        jns pu8_fi

pu_fn_end:
        ; write last character
        call write_char_buffered

        pop r12
        pop rbx

        ret

; read unicode encoded as utf-8 entity
; returns in eax unicode entity number
read_utf8:
        call read_char_buffered

        and eax, CHAR_MASK             ; we want to only have lower 8 bits set

; validate first byte and choose further actions

; check if it's standard ascii character
        cmp ax, UTF8_1BYTE
        jge ru8_if_eof

        ret                            ; if it's single character just end function
; check if it's eof
ru8_if_eof:
        cmp ax, EOF_CHAR
        jne ru8_if_err

        mov eax, EOF                   ; save EOF as 32bit int and return
        ret
; check if it's incorrect first byte
ru8_if_err:
; byte correctness validation
        cmp ax, UTF8_MAX
        jg _error_exit
        cmp ax, UTF8_2BYTES
        jl _error_exit

; this unicode entity is at least 3 bytes long in utf-8 encoding

        push rbx                       ; will be used for further validation
; if it's shortest possible way of writing this character
        push r12                       ; will be used as helper to memorize read character

; check if utf-8 entity is 4 bytes long
        cmp ax, UTF8_4BYTES
        jl ru8_if_3b

        xor al, UTF8_4BYTES
        mov edi, 3                     ; 3 more bytes to read for full character
        mov ebx, UTF8_MIN_VAL_4BYTES

        jmp ru8_fi

ru8_if_3b:                             ; check if utf-8 entity is 3 bytes long

        cmp ax, UTF8_3BYTES
        jl ru8_else

        xor al, UTF8_3BYTES
        mov edi, 2                     ; 2 more bytes to read for full character
        mov ebx, UTF8_MIN_VAL_3BYTES

        jmp ru8_fi

ru8_else:                              ; utf-8 entity is 2 bytes long

        xor al, UTF8_2BYTES
        mov edi, 1                     ; 2 more bytes to read for full character
        mov ebx, UTF8_MIN_VAL_2BYTES

ru8_fi:

        mov ecx, edi
        imul ecx, 6                    ; safe as r8d is at most 3

        shl eax, cl                    ; shift left to accommodate bits from further bytes
        mov r12d, eax                  ; store result

        call read_additional_bytes

        add eax, r12d                  ; add to new bits bits from first byte

        cmp eax, UTF8_MAX_UNICODE_VALUE ; validate that read character isn't too big
        jg _error_exit
        cmp eax, ebx                   ; validate that it's shortest way of writing this character
        jl _error_exit

        pop r12
        pop rbx

        ret

; reads additional utf-8 bytes if utf8 encoding is multibyte
; edi - how many bytes should be read
; in eax returns decoded number from those bytes
read_additional_bytes:
        push rbx                       ; used to store the result
        push r12                       ; will be used as the loop counter

        mov r12d, edi

        xor rbx, rbx                   ; zero result

read_byte:
        call read_char_buffered

; validation
        cmp al, UTF8_BYTE_MASK
        jl _error_exit
        cmp al, UTF8_ADDITIONAL_BYTE_MAX
        jg _error_exit

        xor al, UTF8_BYTE_MASK         ; we only care about 6 lower bits

        movzx r8d, al

; update result
        shl ebx, 6
        add ebx, r8d

        dec r12d
        jnz read_byte

; save result
        mov eax, ebx

        pop r12
        pop rbx

        ret

; writes from out buffer
; updates out_buf_cnt
; in eax is stored actual number of written bytes
write_out_buffer:
        ; actual write syscall
        mov eax, SYS_WRITE
        mov rdi, STDOUT
        mov rsi, out_buf
        mov edx, [out_buf_cnt]

        syscall

        sub DWORD [out_buf_cnt], eax   ; reset out_buf_cnt

        ret

; writes everything from the out buffer
; if encounters error exits program with code 1
flush_out_buff:
        call write_out_buffer
        cmp eax, 0
        js _error_exit                 ; if eax is less than zero it means error
        jne flush_out_buff

        ret

; Writes char to stdout, buffered. Writes in chunks of IN_BUF_SIZE.
; dil - char to write
; returns exit code od syscall
write_char_buffered:
        mov r10d, [out_buf_cnt]
        mov r11, out_buf

        mov [r11 + r10], dil      ; add to out buffer

        inc r10d                       ; increase pointer
        mov [out_buf_cnt], r10d        ; update in memory

        cmp r10d, OUT_BUF_SIZE         ; check if buffer is full

        je write_out_buffer

        ret

; Reads char from stdout, buffered. Reads in chunks of OUT_BUF_SIZE.
; returns read char in al
read_char_buffered:
        mov r10d, [in_buf_reads_cnt]
        mov rsi, in_buf

        cmp r10d, [in_buf_reads_size]  ; check if buffer is not empy
        jnz end_read

        ; actual read syscall
        mov eax, SYS_READ
        mov rdi, STDIN
        mov rdx, IN_BUF_SIZE
        syscall

        mov [in_buf_reads_size], eax   ; save how many characters were read
        xor r10d, r10d                 ; reset reads count

        cmp eax, 0                     ; check if we read anything
        jne end_read

        mov BYTE [rsi], EOF

end_read:

        mov al, [rsi + r10]

        inc r10d
        mov [in_buf_reads_cnt], r10d

        ret

; Parses arguments to int_32t. Overrides argv to: do so.
; edi - argc
; rsi - argv
parse_args:
        lea ecx, [edi - 1]

; validate if at least 1 coefficient was passed
        cmp edi, 2
        js _error_exit

        push r12                       ; loop helper

        mov r12, 1                     ; we start from 1 as 0-th argument is program name

loop_pi:
        mov rdi, [rsi + 8 * r12]       ; i-th argument
        call str_to_unsigned

        mov [rsi + 4 * (r12 - 1)], eax ; set corresponding position in array to int

        inc r12                        ; increase i
        loop loop_pi

        mov rax, rsi

        pop r12

        ret

; calculates diacritization polynomial modulo MOD_VAL
; rdi - x
; rsi - array of coefficients
; rdx - number of coefficients
; calculated value is in rax
calculate_diacritization_polynomial:
        xor rax, rax
        mov rcx, rdx                   ; loop counter

loop_cdp:
; update result using Horner's method
        mul rdi
        mov r8d, [rsi + 4 * (rcx - 1)]
        add rax, r8

; modulo
        cmp rax, MOD_VAL
        js no_modulo_cdp

        mov rdx, 0
        mov r8, MOD_VAL
        div r8
        mov rax, rdx

no_modulo_cdp:

        loop loop_cdp

        ret

; converts string to unsigned 32bit integer modulo MOD_VAL
str_to_unsigned:
        xor eax, eax                   ; zero return register

        mov r8d, MOD_VAL               ; for modulo operation
loop_stu:
; loop condition
        movzx r9d, BYTE [rdi]
        cmp r9d, STR_END
        je exit_stu

; validation
        cmp r9d, ZERO_CHAR
        js _error_exit
        cmp r9d, NINE_CHAR
        ja _error_exit

; return value update
        mov r10d, 10
        mul r10d

        sub r9d, ZERO_CHAR
        add eax, r9d

; modulo
        cmp eax, MOD_VAL ; do not perform expensive modulo if not needed
        js no_modulo_stu

        mov edx, 0
        mov r8d, MOD_VAL
        div r8d
        mov eax, edx
; no modulo needed
no_modulo_stu:

        inc rdi                        ; go to next character

        jmp loop_stu
; loop end
exit_stu:
        ret

; exit with 1 as error code
_error_exit:
        mov edi, 1
        jmp _exit

; exits program with return code as set in rdi
_exit:
        push rdi
        call flush_out_buff            ; flush buffer if anything was in it before system exit
        pop rdi
        mov eax, SYS_EXIT
        syscall
