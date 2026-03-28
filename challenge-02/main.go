package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const validTestFilePath = "./tests/step1/valid.json"
const invalidTestFilePath = "./tests/step1/invalid.json"

func main() {
	validTestFile, err := os.Open(validTestFilePath)
	if err != nil {
		log.Fatal("(Error): ", err)
	}
	defer validTestFile.Close()

	invalidTestFile, err := os.Open(invalidTestFilePath)
	if err != nil {
		log.Fatal("(Error): ", err)
	}
	defer invalidTestFile.Close()

	fmt.Println("Running test for valid one")
	code, err := Parser(validTestFile)
	if err != nil {
		fmt.Println(err)
	}

	if code == 0 {
		fmt.Println("Success")
	}

	fmt.Println("Running test for invalid one")
	code, err = Parser(invalidTestFile)
	if err != nil {
		fmt.Println(err)
	}

	if code == 0 {
		fmt.Println("Success")
	}
}

func Parser(file *os.File) (int, error) {
	stack := []byte{}
	empty := true

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("(Error): ", err)
	}

	for _, byte := range data {
		switch byte {
		case '{':
			stack = append(stack, '{')
			empty = false
			continue

		case '}':
			if len(stack) != 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				continue
			}
			return 1, errors.New("(Error)*: Parse failed")
		}
	}

	if len(stack) == 0 && !empty {
		return 0, nil
	}

	return 1, errors.New("(Error)*: Parse failed")
}
