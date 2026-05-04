// Handles file related work,
// cli related code

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// IpMode type represents
// input mode.
type IpMode int

const (
	// IpModeTerminal represents terminal as input mode.
	IpModeTerminal IpMode = iota
	// IpModeStdin represents stdin as input mode
	IpModeStdIn
)

func getIPMode() (IpMode, error) {

	stdStat, err := os.Stdin.Stat()
	if err != nil {
		return 0, err
	}

	if (stdStat.Mode() & os.ModeCharDevice) == 0 {
		return IpModeStdIn, nil
	}

	return IpModeTerminal, nil
}

func readTerminalInput() (*os.File, error) {
	var err error

	if NoArgs() {
		logrus.Errorf("Cannot read input: %v", ErrNoArgProvided)
		return nil, fmt.Errorf("Cannot read input: %w", ErrNoArgProvided)
	}

	// get file path
	filePath := flag.CommandLine.Arg(0)
	// trim unwanted space
	trimmedFPath := strings.TrimSpace(filePath)
	if trimmedFPath == "" {
		logrus.Errorf("Cannot read input: %v", ErrEmptyFilePath)

		return nil, fmt.Errorf("Cannot read input: %w", ErrEmptyFilePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("Cannot read input: %v", err)
	}

	return f, nil
}

// Format returns ready to print slices of flags and stats.
func Format(fstat *FileInfo, opts *Opts) ([]string, []int) {
	flags := []string{}
	stats := []int{}

	if opts.getByte {
		flags = append(flags, "bytes")
		stats = append(stats, fstat.Bytes)
	}
	if opts.getLine {
		flags = append(flags, "lines")
		stats = append(stats, fstat.Lines)
	}
	if opts.getWord {
		flags = append(flags, "words")
		stats = append(stats, fstat.Words)
	}
	if opts.getRune {
		flags = append(flags, "runes")
		stats = append(stats, fstat.Runes)
	}

	return flags, stats
}

// PrintCol prints the stats with their flag as table.
func PrintCol(flags []string, stats []int) {

	fmt.Println()
	for _, v := range flags {

		fmt.Printf("%-8v", v)
	}

	fmt.Println()
	for _, s := range stats {

		fmt.Printf("%-8d", s)
	}

	fmt.Println()
}

// ReadInput reads input, provided by user.
func ReadInput() (*os.File, error) {

	ipMode, err := getIPMode()
	if err != nil {

		logrus.Errorf("Could not get IP mode: %v", err)
		return nil, err
	}

	file := new(os.File)

	switch ipMode {
	case IpModeTerminal:
		file, err = readTerminalInput()
		if err != nil {
			PrintUsage()
			return nil, err
		}

	case IpModeStdIn:
		file = os.Stdin
		// Todo: add default case
	}

	return file, nil
}
