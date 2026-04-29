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
