package main

import (
	"GoVM/compiler"
	"GoVM/utils"
	"GoVM/vm"
	"os"
)

func main() {
	cliArgs := os.Args

	output := "/home/fs0ciety/Desktop/GoVM/test_data/output.vmb"

	if len(cliArgs) > 1 {
		err := compiler.Compiler(cliArgs[1], output)
		if err != nil {
			utils.CompilationError(err.Error())
			return
		}
		vm.Execution(output)
	} else {
		utils.Error("not found .vms source code file")
	}

}
