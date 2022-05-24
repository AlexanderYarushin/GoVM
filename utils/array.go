package utils

import (
	"errors"
	"fmt"
)

func RemoveEmptyItems(data []string) []string {
	for i := 0; i < len(data); i++ {
		data[i] = ClearSpace(data[i])
		if len(data[i]) == 0 {
			data = append(data[:i], data[i+1:]...)
			i--
		}
	}

	return data
}

func LogArray(data []string) {
	fmt.Print("[")
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i] + ", ")
	}
	fmt.Print("]")
}

func CheckStringDuplicate(data []string) error {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j {
				if data[i] == data[j] {
					return errors.New(Errors["DuplicateLabels"]())
				}
			}
		}
	}

	return nil
}
