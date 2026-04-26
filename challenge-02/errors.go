package main

import "errors"

var (
	// ErrUnexpectedEOF indicates that the input ended before parsing was complete.
	ErrUnexpectedEOF = errors.New("unexpected end of input")

	// ErrIllegalToken indicates that the token is invalid
	// or not recognized by the lexer.
	ErrIllegalToken = errors.New("illegal token")

	// ErrUnknownToken indicates that the token is valid, but appears in an unexpected
	// position according to the grammar.
	ErrUnknownToken = errors.New("unknown token")
)
