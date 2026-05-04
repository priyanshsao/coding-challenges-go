package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

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

func main() {
	// setup logger
	formatLogs()

	myFile := new(FileInfo)
	opts := new(Opts)

	if RegisterAndParse(opts); NoFlags() {
		SetDefaults(opts)
	}

	inStatus, err := os.Stdin.Stat()
	if err != nil {
		logrus.Fatal(err)
	}

	var file *os.File

	if (inStatus.Mode() & os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {

		// stop if no args
		if flag.NArg() == 0 {
			logrus.Errorln("No argument provided")
			flag.Usage()
			os.Exit(1)
		}

		// get the first arg,
		// should be a file path
		filePath := flag.CommandLine.Arg(0)
		// trim unwanted space
		trimmedFPath := strings.TrimSpace(filePath)

		// empty check
		if trimmedFPath != "" {
			file, err = os.Open(filePath)
			if err != nil {
				logrus.Fatal(err)
			}
			defer file.Close()
		} else {
			logrus.Fatal("invalid argument: empty file path.")
		}
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
			logrus.Fatal(err)
		}

		if opts.getByte {
			processBytes(buffer[:n], myFile)
		}
		if opts.getLine {
			processLines(buffer[:n], myFile)
		}
		if opts.getWord {
			processWords(buffer[:n], myFile, &inWord)
		}
		if opts.getRune {
			processRunes(buffer[:n], myFile, &leftOver)
		}
	}

	if opts.getByte {
		logrus.Println(myFile.Bytes)
	}
	if opts.getLine {
		logrus.Println(myFile.Lines)
	}
	if opts.getWord {
		logrus.Println(myFile.Words)
	}
	if opts.getRune {
		logrus.Println(myFile.Runes)
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

func formatLogs() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
		PadLevelText:     true,
	})
}
