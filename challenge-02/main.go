package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type testFile struct {
	name    string
	step    int
	isValid bool
}

var test = []testFile{
	{name: "invalid", step: 1, isValid: false},
	{name: "valid", step: 1, isValid: true},
	{name: "invalid", step: 2, isValid: false},
	{name: "invalid2", step: 2, isValid: false},
	{name: "valid", step: 2, isValid: true},
	{name: "valid2", step: 2, isValid: true},
	{name: "invalid", step: 3, isValid: false},
	{name: "valid", step: 3, isValid: true},
}

func main() {

	for _, t := range test {
		fmt.Printf("\n[step-%d](%s)(%s.json):\n", t.step, getCase(t.isValid), t.name)

		fileData, err := getData("./tests/step" + strconv.Itoa(t.step) + "/" + t.name + ".json")
		if err != nil {
			fmt.Println("[Error]: ", err)
			continue
		}

		p := NewParser(fileData)

		out, err := p.Parse()
		if err != nil {
			fmt.Println("[Fail]: ", err)
		} else {
			fmt.Println("[Pass]: ", out)
		}
	}
}

func getData(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func getCase(b bool) string {
	if b {
		return "valid case"
	}
	return "invalid case"
}
