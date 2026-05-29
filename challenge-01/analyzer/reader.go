package analyzer

import (
	"io"
	"os"
	"strings"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
)

// ipMode represents input mode
type ipMode int

const (
	TERMINAL ipMode = iota
	STD_IN
)

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
			// Todo: add debug.
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
		// Todo: add debug here
		return -1, err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return STD_IN, nil
	}

	return TERMINAL, nil
}

func readTerm(filePath string) (*os.File, error) {

	if trimmedFpath := strings.TrimSpace(filePath); trimmedFpath == "" {
		return nil, ErrEmptyFilePath
	}

	return os.Open(filePath)
}
