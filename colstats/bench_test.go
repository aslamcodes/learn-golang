package main

import (
	"io"
	"path/filepath"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	filenames, err := filepath.Glob("./testdata/benchmark/*.csv")

	if err != nil {
		b.Fatal("the filepaths can't be loaded")
	}

	if len(filenames) == 0 {
		b.Fatal("no files matched")
	}

	for b.Loop() {
		if err := run(filenames, "avg", 1, io.Discard); err != nil {
			b.Fatalf("error performing operation: %v", err)
		}
	}

}
