package config

import "io"

// FileOpts represents a supported file analysis option.
type FileOpts int

// Opts represents the operations requested by the user.
type Opts map[FileOpts]bool

// Result stores the computed values for each option.
type Result map[FileOpts]int

const (
	// BYTES represents byte count operation.
	BYTES FileOpts = iota

	// LINES represents line count operation.
	LINES

	// WORDS represents word count operation.
	WORDS

	// RUNES represents rune count operation.
	RUNES
)

// Config holds the configuration for analyzing an input stream.
type Config struct {
	// Opts contains the selected analysis options.
	Opts Opts

	// Args contains the command-line arguments provided by the user.
	Args []string

	// Reader is the input stream to analyze.
	Reader io.Reader

	// Result stores the computed results for all selected options.
	Result Result
}

// New returns a new config for analyzing input. 
func New() *Config {

	config := &Config{
		Opts:   make(Opts),
		Result: make(Result),
	}

	return config
}

// String value of an operation.
func (opts FileOpts) String() string {

	var str string

	switch opts {
	case BYTES:
		str = "Bytes"
	case LINES:
		str = "Lines"
	case WORDS:
		str = "Words"
	case RUNES:
		str = "Runes"
	}

	return str
}
