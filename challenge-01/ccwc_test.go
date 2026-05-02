package main

import "testing"

func Test_ProcessBytes(t *testing.T) {

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

func Test_ProcessLines(t *testing.T) {

	// Todo: currently processLine only
	// counts \n in future we need to add
	// logic and test for counting lines for files
	// that dont end with '\n' or '\r\n'.

	tests := []struct {
		name     string
		input    string
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

func Test_ProcessWords(t *testing.T) {

	tests := []struct {
		name          string
		input         []string
		expected      int
		expectedState []bool
	}{
		// empty []string{} case
		// is not present, because
		// function's loop doesn't run for that case,
		// so we think there is no need to
		// test that.
		{
			"empty string",
			[]string{""},
			0,
			[]bool{false},
		},
		{
			"normally spaced words",
			[]string{"Hi ", "I am", " a developer"},
			5,
			[]bool{false, true, true},
		},

		// edge cases
		{
			"seperated by escape characters",
			[]string{"Hi,\tI\tam\ta\ndeveloper"},
			5,
			[]bool{true},
		},
		{
			"rune input",
			[]string{"😍", "😍", "😍"},
			1,
			[]bool{true, true, true},
		},
		{
			"incomplete words across buffer",
			[]string{" h", "ey! this i", "s cc", "wc "},
			4,
			[]bool{true, true, true, false},
		},
		{
			"single character in each buffer",
			[]string{"h", "e", "l", "l", "o"},
			1,
			[]bool{true, true, true, true, true},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			inWord := new(bool)
			file := new(FileInfo)

			for i, str := range tt.input {
				processWords([]byte(str), file, inWord)

				if *inWord != tt.expectedState[i] {
					t.Errorf("expected %v but got %v", tt.expectedState[i], *inWord)
				}
			}

			if file.Words != tt.expected {
				t.Errorf("expected %v but got this %v", tt.expected, file.Words)
			}
		})
	}
}
