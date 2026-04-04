package main

import (
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
		log.Println("[Error]: ", err)
		os.Exit(2)
	}
	defer validTestFile.Close()

	invalidTestFile, err := os.Open(invalidTestFilePath)
	if err != nil {
		log.Println("[Error]: ", err)
		os.Exit(2)
	}
	defer invalidTestFile.Close()

	validTestData, err := io.ReadAll(validTestFile)
	if err != nil {
		log.Println("[Error]: ", err)
		os.Exit(2)
	}

	p := NewParser(validTestData)

	output, err := p.Parse()
	if err != nil {
		log.Println("[Error/(step-1)]: ", err)
	} else {
		fmt.Println("[Step-1/(valid Json)]: ", output)
	}

	invalidTestData, err := io.ReadAll(invalidTestFile)
	if err != nil {
		log.Println("[Error]: ", err)
		os.Exit(2)
	}

	p2 := NewParser(invalidTestData)

	output2, err := p2.Parse()
	if err != nil {
		log.Println("[Error]: ", err)
	} else {
		fmt.Println("[Step-1/(invalid json)]: ", output2)
	}
}
