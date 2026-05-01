package main

import "testing"

func Test_processBytes(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty", "", 0},
		{"ascii", "hello", 5},

		// edge cases
		{"escape sequence", "\n\n\t\r", 4},
		{"unicode", "😍😍😍", 12}, //😍 = 4 bytes
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			file := new(FileInfo)

			processBytes([]byte(tt.input), file)

			if file.Bytes != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, file.Bytes)
			}
		})
	}
}

func Test_processLines(t *testing.T) {

	// Todo: currently processLine only
	// counts \n in future we need to add
	// logic and test for counting lines for files
	// that dont end with '\n' or '\r\n'.

	tests := []struct {
		name  string
		input string
		expected int
	}{
		{"line feed", "Hi,\n I am a dev", 1},
		{"CR+LF", "Hi,\r\n I am a dev\r\n", 2},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			file := new(FileInfo)

			processLines([]byte(tt.input), file)

			if file.Lines != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, file.Lines)
			}
		})
	}
}
