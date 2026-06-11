package analyzer

import (
	"io"
	"os"
	"strings"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
	"github.com/sirupsen/logrus"
)

// ipMode represents input mode
type ipMode int

const (
	TERMINAL ipMode = iota
	STD_IN
)

func (ipM ipMode) String() string {
	
	switch ipM {
	case TERMINAL:
		return "Terminal"
	case STD_IN:
		return "Standard input"
	}

	return ""
}

// Read reads the input and stores it in Config.
func Read(c *config.Config) error {

	mode, err := getIpMode()
	if err != nil {
		return err
	}

	var reader io.Reader

	switch mode {
	case TERMINAL:
		if len(c.Args) == 0 {
			logrus.Debugf("unable to read: %s", ErrNoArgProvided)
			return ErrNoArgProvided
		}
		reader, err = readTerm(c.Args[0])
		if err != nil {
			return err
		}

	case STD_IN:
		reader = os.Stdin
	}

	c.Reader = reader

	return nil
}

func getIpMode() (ipMode, error) {

	stat, err := os.Stdin.Stat()
	if err != nil {
		logrus.Debugf("unable to get input mode: %v", err)
		return -1, err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		logrus.Debugf("input mode detected: %v", STD_IN)
		return STD_IN, nil
	}

	logrus.Debugf("input mode detected: %v", TERMINAL)
	return TERMINAL, nil
}

func readTerm(filePath string) (*os.File, error) {

	if trimmedFpath := strings.TrimSpace(filePath); trimmedFpath == "" {
		logrus.Debugf("unable to read from terminal: %s", ErrEmptyFilePath)
		return nil, ErrEmptyFilePath
	}

	return os.Open(filePath)
}
