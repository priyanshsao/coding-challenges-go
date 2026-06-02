package analyzer

import (
	"bytes"
	"io"
	"unicode/utf8"

	"github.com/priyanshsao/coding-challenges-go/challenge-01/config"
)

// Process computes the value of each option.
func Process(c *config.Config) error {

	buffer := make([]byte, 32*1024) //32kB
	leftOver := []byte{}
	inWord := false
	reader := c.Reader

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		for opts, val := range c.Opts {

			if val {

				switch opts {
				case config.BYTES:
					processBytes(buffer[:n], &c.Result)
				case config.LINES:
					processLines(buffer[:n], &c.Result)
				case config.WORDS:
					processWords(buffer[:n], &c.Result, &inWord)
				case config.RUNES:
					processRunes(buffer[:n], &c.Result, &leftOver)
				}
			}
		}
	}

	return nil
}

func processBytes(buffer []byte, result *config.Result) {

	if len(buffer) > 0 {
		(*result)[config.BYTES] += len(buffer)
	}
}

func processLines(buffer []byte, result *config.Result) {

	(*result)[config.LINES] += bytes.Count(buffer, []byte{'\n'})
}

func processWords(buffer []byte, result *config.Result, inWord *bool) {

	for _, b := range buffer {
		if isSpace(b) {
			*inWord = false
		} else {

			if !*inWord {
				(*result)[config.WORDS]++
				*inWord = true
			}
		}
	}
}

func processRunes(buffer []byte, result *config.Result, leftOver *[]byte) {

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

		(*result)[config.RUNES]++
		i += size
	}
}

func isSpace(b byte) bool {
	// simple ASCII check
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}
