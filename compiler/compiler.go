package compiler

import (
	"GoVM/utils"
	"errors"
	"fmt"
	"strings"
)

func Compiler(sourcePath string, targetBinPath string) error {
	buffer, err := utils.ReadFile(sourcePath)

	if err != nil {
		utils.Error(err.Error())
		return nil
	}

	_, compilationErr := translateAsmToBin(buffer)

	if compilationErr != nil {
		utils.CompilationError(compilationErr.Error())
		return nil
	}

	return nil
}

func translateAsmToBin(asm string) ([]byte, error) {
	splitAsmCommands := utils.RemoveEmptyItems(strings.Split(asm, "\n"))

	err := syntaxValidate(splitAsmCommands)

	if err != nil {
		return nil, err
	}

	var bin []byte = translate(splitAsmCommands)

	return bin, nil
}

func translate(commands []string) []byte {
	var bin []byte

	for _, value := range commands {
		value = utils.ClearSpace(value)

		commandSplit := strings.Split(value, " ")

		var command, arg1, arg2 string

		command = commandSplit[0]

		if Commands[command].argsCount >= 1 {
			arg1 = commandSplit[1]

			if Commands[command].argsCount >= 2 {
				arg2 = commandSplit[2]
			}
		}

		fmt.Println(command, arg1, arg2)

	}

	return bin
}

func syntaxValidate(commands []string) error {
	var labels []string

	for _, value := range commands {

		value = utils.ClearSpace(value)

		commandSplit := strings.Split(value, " ")

		var command, arg1, arg2 string

		command = commandSplit[0]

		if Commands[command].argsCount >= 1 {
			arg1 = commandSplit[1]

			if err := isValidRegister(command, arg1); err != nil {
				return err
			}

			if Commands[command].argsCount >= 2 {
				arg2 = commandSplit[2]

				if err := isValidRegister(command, arg2); err != nil {
					return err
				}
			}
		}

		if uint8(len(commandSplit))-1 != Commands[command].argsCount {
			return errors.New("invalid number of arguments in command '" + value + "'")

		}

		if (Commands[command].value == 0) && (!isLoop(command)) {
			return errors.New("invalid command '" + command + "'")
		}

		if isLoop(command) {
			labels = append(labels, command)
		}

	}

	if duplicateLabelError := utils.CheckStringDuplicate(labels); duplicateLabelError != nil {
		return duplicateLabelError
	}

	return nil
}

func isValidRegister(command string, register string) error {
	if !utils.IsInt(register) && Commands[command].value <= 8 {
		if !isRegister(register) {
			return errors.New("invalid operand '" + register + "'")
		}
	}

	return nil
}

func isRegister(register string) bool {
	if Register[register] == 0 {
		return false
	}

	return true
}

func isLoop(command string) bool {
	if strings.Index(command, ":") == -1 {
		return false
	}

	return true
}
