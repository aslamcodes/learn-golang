package main

import "testing"

func TestGetOp(t *testing.T) {
	testcases := []struct {
		opName         string
		expectedResult float64
		data []float64
	}{

	}

	for _, tc := range testcases {
		actual, _ := getOperation(tc.opName)
	}
}
