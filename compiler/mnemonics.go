package compiler

type commandParam struct {
	Value     uint8
	ArgsCount uint8
}

var CommandOption = map[string]uint8{
	"none": 0,
	"rn":   1,
	"rr":   2,
}

var Register = map[string]uint8{
	"ax": 1,
	"bx": 2,
	"cx": 3,
	"dx": 4,
	"ex": 5,
	"fx": 6,
}

var Commands = map[string]commandParam{
	"mov":  {Value: 1, ArgsCount: 2},
	"add":  {Value: 2, ArgsCount: 2},
	"sub":  {Value: 3, ArgsCount: 2},
	"mul":  {Value: 4, ArgsCount: 2},
	"div":  {Value: 5, ArgsCount: 2},
	"cmp":  {Value: 6, ArgsCount: 2},
	"inc":  {Value: 7, ArgsCount: 1},
	"dec":  {Value: 8, ArgsCount: 1},
	"je":   {Value: 9, ArgsCount: 1},
	"jn":   {Value: 10, ArgsCount: 1},
	"jl":   {Value: 11, ArgsCount: 1},
	"jle":  {Value: 12, ArgsCount: 1},
	"jg":   {Value: 13, ArgsCount: 1},
	"jge":  {Value: 14, ArgsCount: 1},
	"call": {Value: 15, ArgsCount: 1},
	"ret":  {Value: 16, ArgsCount: 0},
	"hlt":  {Value: 17, ArgsCount: 0},
	"log":  {Value: 18, ArgsCount: 1},
}

const indexOfValueOperation uint8 = 8
const RegisterNum = 6
