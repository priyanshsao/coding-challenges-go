// Handles flag related work.
// cli related code

package main

import (
	"flag"
	"fmt"
)

type Opts struct {
	// getByte holds the value of '-c' flag.
	getByte bool
	// getLine holds the value of '-l' flag.
	getLine bool
	// getWord holds the value of '-w' flag.
	getWord bool
	// getRune holds the value of '-m' flag.
	getRune bool
}

// NoFlags checks list of flags,
// returns true if list is empty.
func NoFlags() bool {

	return flag.NFlag() == 0
}

// SetDefault sets default flags.
func SetDefaults(opts *Opts) {

	// Todo: add logs here("using default flags")
	opts.getByte = true
	opts.getLine = true
	opts.getWord = true
}

func RegisterAndParse(opts *Opts) {

	// Register 
	flag.BoolVar(&opts.getByte, "c", false, "print total bytes")
	flag.BoolVar(&opts.getLine, "l", false, "print total lines")
	flag.BoolVar(&opts.getWord, "w", false, "print total words")
	flag.BoolVar(&opts.getRune, "m", false, "print total bytes(according to utf-8 encoding)")

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
}