package vm

import (
	"GoVM/compiler"
	"GoVM/utils"
)

func Execution(binPath string) {
	var programBin = utils.ReadBinFile(binPath)

	for {
		var param, operator, arg1, arg2 = programBin[IP], programBin[IP+1], programBin[IP+2], programBin[IP+3]

		if param == compiler.CommandOption["rr"] {
			arg1--
			arg2--
		}

		if param == compiler.CommandOption["rn"] {
			arg1--
		}

		//fmt.Println("ax ", REGISTERS[0])
		//fmt.Println("bx ", REGISTERS[1])
		//fmt.Println("cx ", REGISTERS[2])
		//fmt.Println("-----------------")

		switch operator {
		case compiler.Commands["mov"].Value:
			Mov(param, arg1, arg2)
		case compiler.Commands["add"].Value:
			Add(param, arg1, arg2)
		case compiler.Commands["sub"].Value:
			Sub(param, arg1, arg2)
		case compiler.Commands["mul"].Value:
			Mul(param, arg1, arg2)
		case compiler.Commands["div"].Value:
			Div(param, arg1, arg2)
		case compiler.Commands["cmp"].Value:
			Cmp(param, arg1, arg2)
		case compiler.Commands["inc"].Value:
			Inc(arg1)
		case compiler.Commands["dec"].Value:
			Dec(arg1)
		case compiler.Commands["je"].Value:
			Je(arg1)
		case compiler.Commands["jn"].Value:
			Jn(arg1)
		case compiler.Commands["jl"].Value:
			Jl(arg1)
		case compiler.Commands["jle"].Value:
			Jle(arg1)
		case compiler.Commands["jg"].Value:
			Jg(arg1)
		case compiler.Commands["jge"].Value:
			Jge(arg1)
		case compiler.Commands["call"].Value:
			Call(arg1)
		case compiler.Commands["ret"].Value:
			Ret()
		case compiler.Commands["log"].Value:
			Log(arg1)
		}

		if operator == compiler.Commands["hlt"].Value {
			break
		}

		IP += compiler.CommandBinSize
	}
}
