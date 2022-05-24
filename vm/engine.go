package vm

import (
	"GoVM/compiler"
	"fmt"
)

var STACK [128]uint8
var REGISTERS [compiler.RegisterNum]uint8

var (
	IP uint8 = 0 // instruction pointer
	CR int8  = 0 // compare result
	SP uint8 = 0 // stack pointer
)

func Mov(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		REGISTERS[arg1] = arg2
	case compiler.CommandOption["rr"]:
		REGISTERS[arg1] = REGISTERS[arg2]
	}
}

func Add(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		REGISTERS[arg1] += arg2
	case compiler.CommandOption["rr"]:
		REGISTERS[arg1] += REGISTERS[arg2]
	}
}

func Sub(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		REGISTERS[arg1] -= arg2
	case compiler.CommandOption["rr"]:
		REGISTERS[arg1] -= REGISTERS[arg2]
	}
}

func Mul(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		REGISTERS[arg1] *= arg2
	case compiler.CommandOption["rr"]:
		REGISTERS[arg1] *= REGISTERS[arg2]
	}
}

func Div(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		REGISTERS[arg1] /= arg2
	case compiler.CommandOption["rr"]:
		REGISTERS[arg1] /= REGISTERS[arg2]
	}
}

func Cmp(param uint8, arg1 uint8, arg2 uint8) {
	switch param {
	case compiler.CommandOption["rn"]:
		CR = int8(arg2) - int8(REGISTERS[arg1])
	case compiler.CommandOption["rr"]:
		CR = int8(REGISTERS[arg2]) - int8(REGISTERS[arg1])
	}
}

func Inc(arg1 uint8) {
	REGISTERS[arg1]++
}

func Dec(arg1 uint8) {
	REGISTERS[arg1]--
}

func Je(arg1 uint8) {
	if CR == 0 {
		IP = arg1
	}
}

func Jn(arg1 uint8) {
	if CR != 0 {
		IP = arg1
		IP -= compiler.CommandBinSize
	}
}

func Jl(arg1 uint8) {
	if CR > 0 {
		IP = arg1

	}
}

func Jle(arg1 uint8) {
	if CR >= 0 {
		IP = arg1
	}
}

func Jg(arg1 uint8) {
	if CR < 0 {
		IP = arg1
	}
}

func Jge(arg1 uint8) {
	if CR <= 0 {
		IP = arg1
	}
}

func Call(arg1 uint8) {
	STACK[SP] = IP
	IP = arg1 - compiler.CommandBinSize
	SP++
}

func Ret() {
	SP--
	IP = STACK[SP]
}

func Log(arg1 uint8) {
	fmt.Println(REGISTERS[arg1])
}
