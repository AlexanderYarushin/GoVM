package compiler

import (
	"GoVM/utils"
	"errors"
	"strconv"
	"strings"
)

var CommandBinSize uint8 = 4

/*
	HEX command:
			Operator Operand2
			|        |
		00  00  00  00
		|	    |
		Option  Operand 1

	Option:
		00 - Number Number
		01 - Register Number
		02 - Number Register
		03 - Register Register
*/

func translateAsmToBin(asm string) ([]byte, error) {
	splitAsmCommands := utils.RemoveEmptyItems(strings.Split(asm, "\n"))
	splitAsmCommands = formatCommands(splitAsmCommands)

	err := syntaxValidate(splitAsmCommands)

	if err != nil {
		return nil, err
	}

	return _translate(splitAsmCommands), nil
}

func _translate(commands []string) []byte {
	var bin []byte
	var labelsAddress = make(map[string]uint8)
	var labelIndex uint8 = 0

	for line, value := range commands {
		var command, _, _ string = splitCommandToOp(value)

		if isLoop(command) {
			commands[line] = ""
			labelsAddress[strings.Replace(command, ":", "", -1)] = uint8(line)*CommandBinSize - (labelIndex * CommandBinSize)
			labelIndex++
		}
	}

	commands = utils.RemoveEmptyItems(commands)

	for _, value := range commands {
		var command, arg1, arg2 string = splitCommandToOp(value)
		var option uint8

		if isRegister(arg1) {
			option = CommandOption["rn"]
		}

		if isRegister(arg2) {
			option = CommandOption["nr"]
		}

		if isRegister(arg1) && isRegister(arg2) {
			option = CommandOption["rr"]
		}

		if Commands[command].Value >= 9 && Commands[command].Value <= 15 {
			arg1 = strconv.Itoa(int(labelsAddress[arg1]))
		}

		bin = append(bin, option)
		bin = append(bin, Commands[command].Value)
		appendArgToBin(arg1, &bin)
		appendArgToBin(arg2, &bin)
	}

	return bin
}

func appendArgToBin(arg string, bin *[]byte) {
	if isRegister(arg) {
		*bin = append(*bin, Register[arg])
	} else {
		value, _ := strconv.ParseUint(arg, 10, 8)
		*bin = append(*bin, uint8(value))
	}

}

func formatCommands(commands []string) []string {
	for i, value := range commands {
		commands[i] = utils.ClearSpace(value)
	}

	return commands
}

func syntaxValidate(commands []string) error {
	var labels []string

	for _, value := range commands {
		commandSplit := strings.Split(value, " ")

		var command, arg1, arg2 string = splitCommandToOp(value)

		if err := validateRegister(command, arg1); err != nil {
			return err
		}

		if err := validateRegister(command, arg2); err != nil {
			return err
		}

		if uint8(len(commandSplit))-1 != Commands[command].ArgsCount {
			return errors.New(utils.Errors["InvalidArgNum"](value))
		}

		if (Commands[command].Value == 0) && (!isLoop(command)) {
			return errors.New(utils.Errors["InvalidArgNum"](value))
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

func splitCommandToOp(fullCommand string) (string, string, string) {
	commandSplit := strings.Split(fullCommand, " ")

	var command, arg1, arg2 string

	command = commandSplit[0]

	if Commands[command].ArgsCount >= 1 {
		arg1 = commandSplit[1]

		if Commands[command].ArgsCount >= 2 {
			arg2 = commandSplit[2]
		}
	}

	return command, arg1, arg2
}

func validateRegister(command string, register string) error {
	if isValidRegister(command, register) {
		return nil
	}

	return errors.New(utils.Errors["InvalidOperand"](register))
}

func isValidRegister(command string, register string) bool {
	if !utils.IsInt(register) {
		if (len(register) == 0 || Commands[command].Value > indexOfValueOperation) ||
			Commands[command].Value <= indexOfValueOperation && isRegister(register) {
			return true
		}
		return false
	}
	return true
}

func isRegister(register string) bool {
	return Register[register] != 0
}

func isLoop(command string) bool {
	if strings.Index(command, ":") == -1 {
		return false
	}

	return true
}
