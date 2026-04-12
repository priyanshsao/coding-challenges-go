package main

import (
	"testing"
)

func Test_NewLexer(t *testing.T) {
	input := []byte(`{"Key":"value"}`)
	l := NewLexer(input)

	if l.char != '{' {
		t.Errorf("expected '{' but got %c", l.char)
	}
}
