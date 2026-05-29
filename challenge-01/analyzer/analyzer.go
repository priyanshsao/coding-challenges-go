package analyzer

import (
	"bytes"
	"io"
	"unicode/utf8"

	stdconfig "github.com/priyanshsao/coding-challenges-go/challenge-01/config"
)

// Process computes the value of each option.
func Process(c *stdconfig.Config) error {

	buffer := make([]byte, 32*1024) //32kB
	leftOver := []byte{}
	inWord := false
	reader := config.Reader

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		for opts, val := range config.Opts {

			if val {

				switch opts {
				case stdconfig.BYTES:
					processBytes(buffer[:n], &config.Result)
				case stdconfig.LINES:
					processLines(buffer[:n], &config.Result)
				case stdconfig.WORDS:
					processWords(buffer[:n], &config.Result, &inWord)
				case stdconfig.RUNES:
					processRunes(buffer[:n], &config.Result, &leftOver)
				}
			}
		}
	}

	return nil
}

func processBytes(buffer []byte, result *stdconfig.Result) {

	if len(buffer) > 0 {
		(*result)[stdconfig.BYTES] += len(buffer)
	}
}

func processLines(buffer []byte, result *stdconfig.Result) {

	(*result)[stdconfig.LINES] += bytes.Count(buffer, []byte{'\n'})
}

func processWords(buffer []byte, result *stdconfig.Result, inWord *bool) {

	for _, b := range buffer {
		if isSpace(b) {
			*inWord = false
		} else {

			if !*inWord {
				(*result)[stdconfig.WORDS]++
				*inWord = true
			}
		}
	}
}

func isSpace(b byte) bool {
	// simple ASCII check
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func processRunes(buffer []byte, result *stdconfig.Result, leftOver *[]byte) {

	buffer = append(*leftOver, buffer...)
	*leftOver = (*leftOver)[:0]

	i := 0
	for i < len(buffer) {
		if !utf8.FullRune(buffer[i:]) {
			*leftOver = append(*leftOver, buffer[i:]...)
			break
		}
		// no need to err check
		// as we are sure there is atleast 1 rune ahed
		_, size := utf8.DecodeRune(buffer[i:])

		(*result)[stdconfig.RUNES]++
		i += size
	}
}
