;nasm -f elf64 task3.asm -o task3.o
;gcc -no-pie task3.o -o task3
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
    push rbp
    mov rbp, rsp

    push rbx            ; сохраняем callee-saved
    sub rsp, 8          ; выравнивание стека до 16 байт

    ; printf("Enter N:")
    mov rdi, msg
    xor eax, eax
    call printf

    ; scanf("%d", &n)
    mov rdi, fmt_in
    mov rsi, n
    xor eax, eax
    call scanf

    mov ebx, [n]        ; счетчик оставшихся чисел
    xor r12d, r12d      ; result = 0

loop_start:
    cmp ebx, 0
    je done

    ; scanf("%d", &x)
    mov rdi, fmt_in
    mov rsi, x
    xor eax, eax
    call scanf

    mov eax, [x]

    ; Проверка на полный квадрат
    xor ecx, ecx        ; i = 0

sqrt_loop:
    mov edx, ecx
    imul edx, ecx       ; edx = i*i

    cmp edx, eax
    je is_square
    ja not_square

    inc ecx
    jmp sqrt_loop

is_square:
    inc r12d

not_square:
    dec ebx
    jmp loop_start

done:
    mov rdi, fmt_out
    mov esi, r12d
    xor eax, eax
    call printf

    add rsp, 8
    pop rbx
    pop rbp
    xor eax, eax
    ret

section .note.GNU-stack noalloc noexec nowrite progbits
