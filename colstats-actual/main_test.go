package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestCSVToFloat(t *testing.T) {
	csvData := `a,b,c,nan
1,2,3,a
4,5,6,b
7,8,9,c
`

	testCases := []struct {
		name   string
		col    int
		exp    []float64
		expErr error
		r      io.Reader
	}{
		{
			name:   "Not a number",
			col:    4,
			exp:    nil,
			expErr: ErrNotANumber,
			r:      strings.NewReader(csvData),
		},
		{
			name:   "col 2",
			col:    2,
			exp:    []float64{2, 5, 8},
			expErr: nil,
			r:      strings.NewReader(csvData),
		},
		{
			name:   "col 3",
			col:    3,
			exp:    []float64{3, 6, 9},
			expErr: nil,
			r:      strings.NewReader(csvData),
		},

		{
			name:   "invalid column",
			col:    5,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      strings.NewReader(csvData),
		},
		{
			name:   "invalid column",
			col:    0,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      strings.NewReader(csvData),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			res, err := csvToFloat(tC.r, tC.col)

			if tC.expErr == nil {
				if err != nil {
					t.Errorf("error parsing csv: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error, got nil instead")
				}

				if !errors.Is(err, tC.expErr) {
					t.Errorf("expected %v, got %v", tC.expErr, err)
				}
			}

			if !slices.Equal(res, tC.exp) {
				t.Errorf("expected %v, got %v", tC.exp, res)
			}
		})
	}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		name      string
		col       int
		op        string
		filenames []string
		exp       string
		expErr    error
	}{
		{
			name:      "Golden Path Avg",
			col:       1,
			op:        "avg",
			filenames: []string{"testdata/sample_1.csv", "testdata/sample_2.csv"},
			exp:       "9.642",
			expErr:    nil,
		},
		{
			name:      "Golden Path",
			col:       1,
			op:        "sum",
			filenames: []string{"testdata/sample_1.csv", "testdata/sample_2.csv"},
			exp:       "192.83999999999997",
			expErr:    nil,
		},
		{
			name:      "invalid operation error",
			col:       1,
			op:        "boo",
			filenames: []string{"testdata/sample_2.csv"},
			exp:       "",
			expErr:    ErrInvalidOperation,
		},
		{
			name:      "Not exists",
			col:       1,
			op:        "sum",
			filenames: []string{"testdata/non_existent.csv"},
			exp:       "",
			expErr:    os.ErrNotExist,
		},
		{
			name:      "No Files Error",
			col:       1,
			op:        "sum",
			filenames: []string{},
			exp:       "",
			expErr:    ErrNoFiles,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			out := bytes.NewBuffer(nil)

			opt := options{
				col:       tC.col,
				op:        tC.op,
				filenames: tC.filenames,
			}

			err := run(opt, out)

			if tC.expErr == nil {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error %v, got nil", tC.expErr)
			}

			if !errors.Is(err, tC.expErr) {
				t.Fatalf("expected error %v, got %v", tC.expErr, err)
			}

			fmt.Println(out.String())
			act, _ := strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
			exp, _ := strconv.ParseFloat(tC.exp, 64)

			if exp != act {
				t.Errorf("expected %f, got %f", exp, act)
			}

		})
	}
}
