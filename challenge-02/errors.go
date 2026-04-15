package main

import "errors"

var (
	ErrUnexpectedEOF = errors.New("unexpected end of input")
	ErrIllegalToken  = errors.New("illegal token")
	ErrUnknownToken  = errors.New("unknown token")
)
