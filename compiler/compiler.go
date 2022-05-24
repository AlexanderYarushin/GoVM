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

	//for i := 0; i < len(bin); i += 4 {
	//	fmt.Println(bin[i], bin[i+1], bin[i+2], bin[i+3])
	//}

	utils.WriteFile(targetBinPath, bin)

	return nil
}
