package main

import (
	"slices"
	"strings"
	"testing"
)

func TestCsv2Float(t *testing.T) {
	tests := []struct {
		name    string
		csv     string
		column  int
		want    []float64
		wantErr bool
	}{
		{"ok", "a,b\n50.4,1\n139.12,2\n11.123,3\n", 1,
			[]float64{50.4, 139.12, 11.123}, false},

		{"column zero", "a,b\n1,2\n", 0, nil, true},

		{"empty csv", "a,b\n", 1, nil, true},

		{"column out of range", "a\n1\n2\n", 2, nil, true},

		{"bad float", "a,b\n50.4,x\n", 2, nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.csv)
			got, err := csv2float(r, tc.column)

			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(got, tc.want) {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestRun(t *testing.T) {
}
