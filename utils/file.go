package utils

import (
	"bufio"
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) (string, error) {
	buffer, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func ReadBinFile(path string) []byte {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fi, _ := os.Stat(path)

	reader := bufio.NewReader(f)

	buffer := make([]byte, fi.Size())

	reader.Read(buffer)

	return buffer
}

func WriteFile(path string, buffer []byte) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	writeErr := binary.Write(file, binary.LittleEndian, buffer)

	if writeErr != nil {
		return writeErr
	}

	closeErr := file.Close()

	if closeErr != nil {
		return closeErr
	}

	return nil
}
