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

	reader := csv.NewReader(r)

	csvData, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("cannot read csv: %w", err)
	}

	dst := []float64{}

	for i, row := range csvData {
		// skip the heading row
		if i == 0 {
			continue
		}

		if len(row) < column {
			return []float64{}, fmt.Errorf("column number out of bounds: %w", ErrInvalidColumn)
		}

		f, err := strconv.ParseFloat(row[column], 64)

		if err != nil {
			return []float64{}, fmt.Errorf("%w: %s", ErrNotANumber, err)
		}

		dst = append(dst, f)
	}

	return dst, nil
}
