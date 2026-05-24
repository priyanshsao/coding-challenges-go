package analyzer

import (
	"io"
	"os"
	"strings"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
)

type ipMode int

const (
	TERMINAL ipMode = iota
	STD_IN
)

// Read reads the input
func Read(config *config.Config) error {

	mode, err := getIpMode()
	if err != nil {
		return err
	}

	var reader io.Reader

	switch mode {
	case TERMINAL:
		if len(config.Args) == 0 {
			// Todo: add debug.
			return ErrNoArgProvided
		}
		reader, err = readTerm(config.Args[0])
		if err != nil {
			return err
		}

	case STD_IN:
		reader = os.Stdin
	}

	config.Reader = reader

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

func readTerm(filePath string) (io.Reader, error) {

	if trimmedFpath := strings.TrimSpace(filePath); trimmedFpath == "" {
		return nil, ErrEmptyFilePath
	}

	return os.Open(filePath)
}
