package main

import (
	"flag"
	"fmt"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/analyzer"
	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
	"github.com/sirupsen/logrus"
)

func main() {

	setLogger()

	c := config.New()
	opts := c.Opts

	// Register
	byteFlag := flag.Bool("c", false, "print total bytes")
	lineFlag := flag.Bool("l", false, "print total lines")
	wordFlag := flag.Bool("w", false, "print total words")
	runeFlag := flag.Bool("m", false, "print total bytes(according to utf-8 encoding)")

	// Set usage
	flag.Usage = func() {

		fmt.Print("\nUsage: ccwc [flag] [file_path]\n")
		fmt.Print("\nDefault flags: -c -w -l\n")
		fmt.Print("\nFlags:\n")
		flag.PrintDefaults()
		fmt.Println() // for readability
	}

	// parses the flags and fills the variables,
	// Should be called before flags are accessed
	// by program
	flag.Parse()

	c.Args = flag.CommandLine.Args()

	if flag.NFlag() == 0 {
		// Todo: add debug logs

		opts[config.BYTES] = true
		opts[config.LINES] = true
		opts[config.WORDS] = true
	} else {

		opts[config.BYTES] = *byteFlag
		opts[config.LINES] = *lineFlag
		opts[config.WORDS] = *wordFlag
		opts[config.RUNES] = *runeFlag
	}

	if err := analyzer.Read(c); err != nil {
		// Todo: log error
		logrus.Errorf("unable to read input: %v", err)
		return
	}

	if err := analyzer.Process(c); err != nil {
		// Todo: log error
		logrus.Errorf("unable to process input: %v", err)
		return
	}

	fmt.Println(c.Result)
}

func setLogger() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableLevelTruncation: true,
	})
}
