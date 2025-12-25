package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	col := flag.Int("col", 0, "the column number to perform the operation (starts at 1)")
	op := flag.String("op", "", "operation to perform on the column")

	flag.Parse()

	filenames := flag.Args()

	opts := options{
		col:       *col,
		op:        *op,
		filenames: filenames,
	}

	if err := run(opts, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type options struct {
	col       int
	op        string
	filenames []string
}

func run(opts options, out io.Writer) error {
	if len(opts.filenames) == 0 {
		return ErrNoFiles
	}

	var op statFunc

	switch opts.op {
	case "avg":
		op = avg
	case "sum":
		op = sum
	default:
		return fmt.Errorf("invalid operation provided")
	}

	consolidated := []float64{}

	for _, f := range opts.filenames {
		file, err := os.Open(f)

		if err != nil {
			return fmt.Errorf("cannot read file: %w", err)
		}

		data, err := csvToFloat(file, opts.col)

		if err != nil {
			return fmt.Errorf("error calculating column %d on file %s: %w", opts.col, file.Name(), err)
		}

		consolidated = append(consolidated, data...)
	}

	_, err := fmt.Fprintln(out, op(consolidated))

	return err
}
