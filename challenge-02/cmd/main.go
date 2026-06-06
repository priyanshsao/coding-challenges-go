package main

import (
	"fmt"
	"io"
	"os"

	"github.com/priyanshsao/coding-challenges-go/challenge-02/json"
	"github.com/sirupsen/logrus"
)

func main() {

	setLogger()

	path := "./tests/step2/invalid2.json"
	file, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer file.Close()

	fileData, _ := io.ReadAll(file)

	parser := json.NewParser(fileData)

	res, err := parser.Parse()
	if err != nil {
		logrus.Errorf("unable to parse: %v", err)
		return
	}

	fmt.Println(res)
}

func setLogger() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	// set log level to debug.
	logrus.SetLevel(logrus.DebugLevel)
}
