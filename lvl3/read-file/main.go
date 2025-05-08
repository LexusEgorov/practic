package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const MAX_FILE_SIZE = 1048576

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("file path is required!")
	}

	fmt.Println(ReadFile(args[1]))
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return "", err
	}

	if info.Size() > MAX_FILE_SIZE {
		return "", errors.New("file is too large")
	}

	data := make([]byte, 16)

	for {
		_, err = file.Read(data)

		if err == io.EOF {
			break
		}
	}

	return string(data), nil
}
