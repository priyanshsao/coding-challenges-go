package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

type FileInfo struct {
	Bytes int
	Lines int
	Words int
	Runes int
}

func main() {
	var getBytes bool
	var getLines bool
	var getWords bool
	var getRunes bool

	myFile := &FileInfo{}

	flag.BoolVar(&getBytes, "c", false, "print total bytes")
	flag.BoolVar(&getLines, "l", false, "print total lines")
	flag.BoolVar(&getWords, "w", false, "print total words")
	flag.BoolVar(&getRunes, "m", false, "print total bytes(according to utf-8 encoding)")

	flag.Parse()

	filePath := flag.CommandLine.Arg(0)

	inStatus, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File

	if (inStatus.Mode() & os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Println("(error): ", err)
		}
		defer file.Close()
	}

	buffer := make([]byte, 32*1024)
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
	file.Lines += bytes.Count(buffer, []byte{'\n'})
}

func processWords(buffer []byte, file *FileInfo, inWord *bool) {
	for _, b := range buffer {
		if unicode.IsSpace(rune(b)) {
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
		// need to add error check for utf8.RuneError
		_, size := utf8.DecodeRune(buffer[i:])

		file.Runes++
		i += size
	}
}
