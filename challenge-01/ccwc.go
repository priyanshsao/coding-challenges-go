package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func formatLogs() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
		PadLevelText:     true,
	})
}

func main() {

	formatLogs()

	opts := new(Opts)

	if RegisterAndParse(opts); NoFlags() {
		SetDefaults(opts)
	}

	file, err := ReadInput()
	if err != nil {
		os.Exit(1)
	}

	fstat, err := Compute(file, opts)
	if err != nil {

		logrus.Errorf("Unable to compute file stats: %v", err)
		os.Exit(1)
	}

	flags, stat := Format(fstat, opts)
	PrintCol(flags, stat)
}
