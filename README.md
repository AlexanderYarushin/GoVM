#### This is a register-based virtual machine written in Go
##### Registers:
    ax, bx, cx, dx, ex, fx
##### Commands
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
