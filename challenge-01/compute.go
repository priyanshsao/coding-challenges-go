// handles the main business logic

package main

import (
	"bytes"
	"io"
	"os"
	"unicode/utf8"
)

// FileInfo holds stats of a file.
type FileInfo struct {
	// count of bytes
	Bytes int
	// count of lines
	Lines int
	// count of words
	Words int
	// count of unicode chars
	Runes int
}

func isSpace(b byte) bool {
	// simple ASCII check
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func processBytes(buffer []byte, file *FileInfo) {

	if len(buffer) > 0 {
		file.Bytes += len(buffer)
	}
}

func processLines(buffer []byte, file *FileInfo) {
	// counts \n in buffer
	file.Lines += bytes.Count(buffer, []byte{'\n'})
}

func processWords(buffer []byte, file *FileInfo, inWord *bool) {

	for _, b := range buffer {
		if isSpace(b) {
			*inWord = false
		} else {

			if !*inWord {
				file.Words++
				*inWord = true
			}
		}
	}
}

func processRunes(buffer []byte, file *FileInfo, leftOver *[]byte) {

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

		file.Runes++
		i += size
	}
}

// Compute computes the file stats according to the flags provided.
func Compute(f *os.File, opts *Opts) (*FileInfo, error) {

	fstat := new(FileInfo)
	buffer := make([]byte, 32*1024) //32kB
	leftOver := []byte{}
	inWord := false

	for {
		n, err := f.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if opts.getByte {
			processBytes(buffer[:n], fstat)
		}
		if opts.getLine {
			processLines(buffer[:n], fstat)
		}
		if opts.getWord {
			processWords(buffer[:n], fstat, &inWord)
		}
		if opts.getRune {
			processRunes(buffer[:n], fstat, &leftOver)
		}
	}

	return fstat, nil
}
