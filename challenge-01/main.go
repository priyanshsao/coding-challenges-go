package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
