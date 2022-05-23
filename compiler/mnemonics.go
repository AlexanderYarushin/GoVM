package compiler

type commandParam struct {
	value     uint8
	argsCount uint8
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
	"mov":  {value: 1, argsCount: 2},
	"add":  {value: 2, argsCount: 2},
	"sub":  {value: 3, argsCount: 2},
	"mul":  {value: 4, argsCount: 2},
	"div":  {value: 5, argsCount: 2},
	"cmp":  {value: 6, argsCount: 2},
	"inc":  {value: 7, argsCount: 1},
	"dec":  {value: 8, argsCount: 1},
	"je":   {value: 9, argsCount: 1},
	"jn":   {value: 10, argsCount: 1},
	"jl":   {value: 11, argsCount: 1},
	"jle":  {value: 12, argsCount: 1},
	"jg":   {value: 13, argsCount: 1},
	"jge":  {value: 14, argsCount: 1},
	"call": {value: 15, argsCount: 1},
	"ret":  {value: 16, argsCount: 0},
	"hlt":  {value: 17, argsCount: 0},
}
