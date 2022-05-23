package utils

import (
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
