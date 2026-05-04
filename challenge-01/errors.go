// Handles error related work

package main

import "errors"

var (
	// ErrEmptyFilePath indicates, provided file path is empty.
	ErrEmptyFilePath error = errors.New("empty file path.")
	// ErrNoArgProvided indicates, no argument is provided.
	ErrNoArgProvided error = errors.New("no argument provided.")
)
