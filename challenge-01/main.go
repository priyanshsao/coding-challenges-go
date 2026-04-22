package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

type FileInfo struct {
	// count of bytes
	Bytes int
	// count of lines
	Lines int
	// count of words
	Words int
	// count of utf-8 chars
	Runes int
}

func main() {
	// variables to store flag value
	var getBytes bool
	var getLines bool
	var getWords bool
	var getRunes bool

	myFile := new(FileInfo)

	// define flags
	flag.BoolVar(&getBytes, "c", false, "print total bytes")
	flag.BoolVar(&getLines, "l", false, "print total lines")
	flag.BoolVar(&getWords, "w", false, "print total words")
	flag.BoolVar(&getRunes, "m", false, "print total bytes(according to utf-8 encoding)")

	// parses the flags and fills the variables,
	// Should be called before flags are accessed
	// by program
	flag.Parse()

	inStatus, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File

	if (inStatus.Mode() & os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {

		// get the first non-flag arg,
		// generally a file path
		filePath := flag.CommandLine.Arg(0)

		// empty check path = ""
		if len(filePath) != 0 {
			file, err = os.Open(filePath)
			if err != nil {
				log.Println("(error): ", err)
			}
			defer file.Close()
		}
		log.Fatal("(error): empty file path")
	}

	buffer := make([]byte, 32*1024) //32kB
	leftOver := []byte{}
	inWord := false

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("(error): ", err)
		}

		if getBytes {
			processBytes(buffer[:n], myFile)
		}
		if getLines {
			processLines(buffer[:n], myFile)
		}
		if getWords {
			processWords(buffer[:n], myFile, &inWord)
		}
		if getRunes {
			processRunes(buffer[:n], myFile, &leftOver)
		}
	}

	if getBytes {
		fmt.Print(myFile.Bytes, " ")
	}
	if getLines {
		fmt.Print(myFile.Lines, " ")
	}
	if getWords {
		fmt.Print(myFile.Words, " ")
	}
	if getRunes {
		fmt.Print(myFile.Runes, " ")
	}
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

func isSpace(b byte) bool {
	// simple ASCII check
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}
