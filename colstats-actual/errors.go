package main

import "errors"

var (
	ErrNotANumber = errors.New("Data is not numeric")
	ErrInvalidColumn = errors.New("Invalid column number")
	ErrNoFiles = errors.New("No input files")
	ErrInvalidOperation = errors.New("Invalid Operation")
)
