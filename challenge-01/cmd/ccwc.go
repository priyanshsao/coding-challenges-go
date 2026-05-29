package main

import (
	"flag"
	"fmt"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/analyzer"
	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
	stdconfig "github.com/priyanshsao/coding-challenges-go/challenge-01/config"
	"github.com/sirupsen/logrus"
)

func main() {

	setLogger()

	config := config.New()
	opts := config.Opts

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

	config.Args = flag.CommandLine.Args()

	if flag.NFlag() == 0 {

		opts[stdconfig.BYTES] = true
		opts[stdconfig.LINES] = true
		opts[stdconfig.WORDS] = true
	} else {

		opts[stdconfig.BYTES] = *byteFlag
		opts[stdconfig.LINES] = *lineFlag
		opts[stdconfig.WORDS] = *wordFlag
		opts[stdconfig.RUNES] = *runeFlag
	}

	if err := analyzer.Read(config); err != nil {
		// log error
		return
	}

	if err := analyzer.Process(config); err != nil {
		// log error
		return
	}

	fmt.Println(config.Result)
}

func setLogger() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableLevelTruncation: true,
	})
}
