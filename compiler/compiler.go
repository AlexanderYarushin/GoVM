package compiler

import (
	"GoVM/utils"
)

func Compiler(sourcePath string, targetBinPath string) error {
	buffer, err := utils.ReadFile(sourcePath)

	if err != nil {
		utils.Error(err.Error())
		return nil
	}

	bin, compilationErr := translateAsmToBin(buffer)

	if compilationErr != nil {
		return compilationErr
	}

	utils.WriteFile(targetBinPath, bin)

	return nil
}
