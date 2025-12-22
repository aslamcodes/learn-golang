package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	op := flag.String("op", "", "The operation to perform")
	col := flag.Int("col", 0, "The column to perform the operation (starting from 1)")

	flag.Parse()

	if err := run(flag.Args(), *op, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(files []string, op string, col int, writer io.Writer) error {
	operation, err := getOperation(op)

	if err != nil {
		return err
	}

	wholeData := make([]float64, len(files))

	if len(files) == 0 {
		return fmt.Errorf("No Files are specified")
	}

	for _, file := range files {
		file, err := os.Open(file)
		defer file.Close()

		if err != nil {
			return err
		}

		fileData, err := csv2float(file, col)

		if err != nil {
			return err
		}

		wholeData = append(wholeData, fileData...)
	}

	_, err = fmt.Fprintln(writer, operation(wholeData))

	return err
}

func csv2float(file io.Reader, column int) ([]float64, error) {
	if column == 0 {
		return nil, fmt.Errorf("No column specified")
	}
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("Empty CSV File with no records")
	}

	if len(records[0]) < column {
		return nil, fmt.Errorf("The csv files does not have enough columns")
	}

	data := make([]float64, len(records)-1)
	for i, record := range records {
		if i == 0 {
			continue
		}

		datum, err := strconv.ParseFloat(record[column-1], 64)

		if err != nil {
			return nil, fmt.Errorf("%w : %s", err, record[column-1])
		}

		data[i-1] = datum
	}

	return data, nil
}
