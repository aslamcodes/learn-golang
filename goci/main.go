package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	project := flag.String("p", "", "project directory")

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
	args := []string{"build", ".", "errors"}

	cmd := exec.Command("go", args...)

	cmd.Dir = project

	if err := cmd.Run(); err != nil {
		return &stepErr{
			step: "go build",
			msg: "go build failed",
			cause: err,
		}
	}

	_, err := fmt.Fprintln(out, "go build success")

	return err
}
