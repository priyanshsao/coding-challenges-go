package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/priyanshsao/coding-challenges-go/challenge-02/json"
	"github.com/sirupsen/logrus"
)

func main() {

	const testDir = "./tests"

	setLogger()

	dirs, err := processDir(testDir)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, dir := range dirs {

		path := filepath.Join(testDir, dir.Name())
		files, err := processDir(path)
		if err != nil {
			logrus.Error(err)
			return
		}

		process(files, path)
	}
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

func processDir(path string) ([]os.DirEntry, error) {

	dirs, err := os.ReadDir(path)
	if err != nil {

		return dirs, err
	}

	return dirs, nil
}

func process(files []os.DirEntry, root string) {

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		path := filepath.Join(root, file.Name())

		fmt.Printf("\n-----------------[%v]-----------------\n\n", path)

		file, err := os.Open(path)
		if err != nil {
			logrus.Error(err)
			continue
		}
		defer file.Close()

		fileData, _ := io.ReadAll(file)

		parser := json.NewParser(fileData)

		res, err := parser.Parse()
		if err != nil {
			logrus.Errorf("unable to parse: %v", err)
			continue
		}

		logrus.Info(res)
	}
}
