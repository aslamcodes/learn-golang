package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	goldenFile string = "./testdata/test.md.html"
	inputFile  string = "./testdata/test.md"
	outputFile string = "test.md"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)

	if err != nil {
		t.Fatalf("readfile: cannot read input file %s. error: %v", inputFile, err)
		os.Exit(1)
	}

	result := parseContent(input)

	expected, err := os.ReadFile(goldenFile)

	if err != nil {
		t.Fatalf("readfile: cannot read golden file %s. error: %v", goldenFile, err)
		os.Exit(1)
	}

	if !bytes.Equal(expected, result) {
		t.Fatalf("mismatch:\nexpected:\n%s\ngot:\n%s", expected, result)
	}

}

func TestRun(t *testing.T) {
	mockStdOut := bytes.Buffer{}

	err := run(inputFile, &mockStdOut, true)

	outName := strings.TrimSpace(mockStdOut.String())

	if err != nil {
		t.Fatalf("error while running the run wrapper %v", err)
	}

	outputContent, err := os.ReadFile(outName)

	if err != nil {
		t.Fatalf("error while reading the outputcontent %v", err)
	}

	expectedContent, err := os.ReadFile(goldenFile)

	if err != nil {
		t.Fatalf("error while reading the expected content %v", err)
	}

	if !bytes.Equal(expectedContent, outputContent) {
		t.Fatalf("mismatch:\nexpected:\n%s\ngot:\n%s", expectedContent, outputContent)
	}

	os.Remove(outName)
}
