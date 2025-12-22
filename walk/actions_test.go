package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{
			name:     "FilterNoExtensions",
			file:     "testdata/dir.log",
			ext:      "",
			minSize:  0,
			expected: false,
		},
		{
			name:     "FilterExtentionMatch",
			file:     "testdata/dir.log",
			ext:      ".log",
			minSize:  0,
			expected: false,
		},
		{
			name:     "FilterExtensionNoMatch",
			file:     "testdata/dir.log",
			ext:      ".txt",
			minSize:  0,
			expected: true,
		},
		{
			name:     "FilterSizeMatch",
			file:     "testdata/dir.log",
			ext:      "",
			minSize:  12,
			expected: false,
		},
		{
			name:     "FilterSizeNoMatch",
			file:     "testdata/dir.log",
			ext:      "",
			minSize:  20,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Stat(tc.file)

			if err != nil {
				t.Fatal(err)
			}

			actual := filterOut(tc.file, tc.ext, tc.minSize, file)

			if tc.expected  != actual {
				t.Errorf("Expected %t, got %t", tc.expected, actual)
			}
		})
	}

}
