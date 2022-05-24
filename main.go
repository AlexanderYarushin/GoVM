package main

import (
	"GoVM/compiler"
	"GoVM/utils"
	"GoVM/vm"
	"fmt"
	"os"
	"time"
)

func main() {
	cliArgs := os.Args

	output := "/home/fs0ciety/Desktop/GoVM/test_data/output.vmb"

	if len(cliArgs) > 1 {
		start := time.Now()

		if err := compiler.Compiler(cliArgs[1], output); err != nil {
			utils.CompilationError(err.Error())
			return
		}

		duration := time.Since(start)

		LogCompilationInfo(output, duration)

		vm.Execution(output)
	} else {
		utils.Error(utils.Errors["SourceNotFound"]())
	}
}

func LogCompilationInfo(binaryFile string, duration time.Duration) {
	fi, _ := os.Stat(binaryFile)

	fmt.Println("Compilation was successful:", duration)
	fmt.Println("Binary size:", fi.Size(), "byte")
	fmt.Println("Output: ")
}
