;только на линуксе т.к. это nasm
;nasm -f elf64 task3.asm -o task3.o
;gcc task3.o -o task3
;./task3
section .data
    msg db "Enter N:", 10, 0
    fmt_in db "%d", 0
    fmt_out db "Result: %d", 10, 0

section .bss
    n resd 1
    x resd 1

section .text
    extern printf, scanf
    global main

main:
    ; print "Enter N:"
    mov rdi, msg
    xor rax, rax
    call printf

    ; read n
    mov rdi, fmt_in
    mov rsi, n
    xor rax, rax
    call scanf

    mov ecx, [n]     ; counter
    xor ebx, ebx     ; count = 0

loop_start:
    cmp ecx, 0
    je done

    ; read x
    mov rdi, fmt_in
    mov rsi, x
    xor rax, rax
    call scanf

    mov eax, [x]

    ; root = sqrt(x) (очень грубо через цикл)
    xor edx, edx
sqrt_loop:
    imul edx, edx
    cmp edx, eax
    je is_square
    ja not_square
    inc edx
    jmp sqrt_loop

is_square:
    inc ebx

not_square:
    dec ecx
    jmp loop_start

done:
    mov rdi, fmt_out
    mov rsi, rbx
    xor rax, rax
    call printf

    ret