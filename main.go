package main

import (
	"GoVM/compiler"
	"GoVM/utils"
	"os"
)

func main() {
	cliArgs := os.Args

	if len(cliArgs) > 1 {
		compiler.Compiler(cliArgs[1], "output.vmb")
	} else {
		utils.Error("not found .vms source code file")
	}

}
