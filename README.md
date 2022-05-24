#### This is a register-based virtual machine written in Go
##### Virtual machine supports the call subprogram
##### General registers:
    ax, bx, cx, dx, ex, fx
##### Service registers:
    ip - instruction pointer
	cr - compare result
	sp - stack pointer
##### Commands:
    mov - moving values
	add - adding values
	sub - substruction values
	mul - multiplication values
	div - division values
	cmp - comparison values
	inc - increment value
	dec - decrement value
	je - goto if equal
	jn - goto if not equal
	jl - goto if less
	jle - goto if less or equal
	jg - goto if great
	jge - goto if great or equal
	call - move to subprogram address
	ret - return from subprogram
	hlt - exit
	log - value output

### Example #1 - calculate factorial(5)
![](/factorial.png)

### Example #2 - calculate and output fibonacci numbers(12)
![](/fibonacci.png)
