package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
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
		return fmt.Errorf("invalid operation provided: %w", ErrInvalidOperation)
	}

	consolidated := []float64{}

	resCh := make(chan []float64)
	errCh := make(chan error)
	doneCh := make(chan struct{})
	filesCh := make(chan string)

	go func() {
		defer close(filesCh)
		for _, f := range opts.filenames {
			filesCh <- f
		}
	}()

	wg := sync.WaitGroup{}

	for range runtime.NumCPU() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fname := range filesCh {
				file, err := os.Open(fname)

				if err != nil {
					errCh <- fmt.Errorf("cannot read file: %w", err)
					return
				}

				data, err := csvToFloat(file, opts.col)

				if err != nil {
					errCh <- fmt.Errorf("error calculating column %d on file %s: %w", opts.col, file.Name(), err)
					return
				}

				if err := file.Close(); err != nil {
					errCh <- err
				}

				resCh <- data
			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return err
		case data := <-resCh:
			consolidated = append(consolidated, data...)
		case <-doneCh:
			_, err := fmt.Fprintln(out, op(consolidated))
			return err
		}
	}

}
