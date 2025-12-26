package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// encoding/csv - Provides methods to work with csv files
// strconv - to convert string to numbers read from the csv columns
// io - For io.Reader interface, can be used to abstract the logic

// functions with same signature, can be defined its type
type statFunc func(data []float64) float64

func sum(data []float64) float64 {
	// var total float64 = 0
	total := 0.0

	for _, datum := range data {
		total += datum
	}

	return total
}

func avg(data []float64) float64 {

	l := float64(len(data))
	if l == 0 {
		return 0
	}

	return sum(data) / l
}

func csvToFloat(r io.Reader, column int) ([]float64, error) {
	column--

	if column < 0 {
		return nil, ErrInvalidColumn
	}

	reader := csv.NewReader(r)
	reader.ReuseRecord = true

	header_row, err := reader.Read()

	if err != nil {
		return nil, fmt.Errorf("error reading from file: %w", err)
	}

	if len(header_row) <= column {
		return nil, fmt.Errorf("column no greater than list of columns: %w", ErrInvalidColumn)
	}

	dst := []float64{}

	for i := 0; ; i++ {

		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading from file: %w", err)
		}

		raw := row[column]

		val, err := strconv.ParseFloat(raw, 64)

		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotANumber, err)
		}

		dst = append(dst, val)
	}

	return dst, nil
}
