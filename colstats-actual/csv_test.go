package main

import "testing"

func TestOperations(t *testing.T) {
	data := [][]float64{
		{},
		{1, 2, 3},
		{1, -3},
		{2.5, 3.5, 4.5},
	}

	testCases := []struct {
		name string
		op   statFunc
		exp  []float64
	}{
		{
			name: "sum",
			op:   sum,
			exp:  []float64{0, 6, -2, 10.5},
		},
		{
			name: "avg",
			op:   avg,
			exp:  []float64{0, 2, -1, 3.5},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			for i, exp := range tC.exp {
				res := tC.op(data[i])

				if res != exp {
					t.Errorf("expected %g, got %g instead", exp, res)
				}

			}
		})
	}

}
