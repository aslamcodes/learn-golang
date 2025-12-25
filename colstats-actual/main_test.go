package main

import (
	"errors"
	"io"
	"slices"
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
	}
}
