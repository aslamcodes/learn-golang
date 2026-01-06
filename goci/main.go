package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type executor interface {
	execute() (string, error)
}

func main() {
	project := flag.String("p", "", "project directory")

	flag.Parse()

	if err := run(*project, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(project string, out io.Writer) error {
	if project == "" {
		return fmt.Errorf("project directory is required")
	}

	// build command won't create binaries when building multiple packages at the same time.
	// Single package → build + link → write executable; Multiple packages → compile-only → cache results
	// go build ./... must be safe to run anywhere without littering directories with binaries
	// A binary maps to exactly one main package
	// go build validates and caches by default; it only produces files when explicitly obvious.
	pipeline := make([]executor, 3)

	pipeline[0] = newStep("go build", project, "go build SUCCESS", "go", []string{"build", ".", "errors"})

	pipeline[1] = newStep("go test", project, "go test OK", "go", []string{"test", "-v"})

	pipeline[2] = NewExecutionStep("go fmt", project, "go format OK", "gofmt", []string{"-l", "."})

	for _, s := range pipeline {
		success_msg, err := s.execute()

		if err != nil {
			return err
		}

		fmt.Fprintln(out, success_msg)
	}

	return nil
}
