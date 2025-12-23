package main

import (
	"testing"
)

func TestGetOp(t *testing.T) {
	testcases := []struct {
		OpName         string
		ExpectedResult float64
		Data           []float64
	}{

		{
			OpName:         "avg",
			ExpectedResult: 0,
			Data:           []float64{},
		},
		{
			OpName:         "avg",
			ExpectedResult: 2,
			Data:           []float64{1, 2, 3},
		},
		{
			OpName:         "avg",
			ExpectedResult: -1,
			Data:           []float64{-2, 0},
		},
		{
			OpName:         "avg",
			ExpectedResult: 2.5,
			Data:           []float64{2, 3},
		},
		{
			OpName:         "sum",
			ExpectedResult: 0,
			Data:           []float64{},
		},
		{
			OpName:         "sum",
			ExpectedResult: 6,
			Data:           []float64{1, 2, 3},
		},
		{
			OpName:         "sum",
			ExpectedResult: -2,
			Data:           []float64{1, -3},
		},
		{
			OpName:         "sum",
			ExpectedResult: 10.5,
			Data:           []float64{2.5, 3.5, 4.5},
		},
	}

	for _, tC := range testcases {

		fn, err := getOperation(tC.OpName)

		if err != nil {
			t.Errorf("error getting operation: %v", err)
		}

		a := fn(tC.Data)

		if a != tC.ExpectedResult {
			t.Errorf("expected %f, got %f", tC.ExpectedResult, a)
		}
	}
}

func TestSum(t *testing.T) {
	testcases := []struct {
		OpName         string
		ExpectedResult float64
		Data           []float64
	}{
		{
			ExpectedResult: 0,
			Data:           []float64{},
		},
		{
			ExpectedResult: 6,
			Data:           []float64{1, 2, 3},
		},
		{
			ExpectedResult: -2,
			Data:           []float64{1, -3},
		},
		{
			ExpectedResult: 10.5,
			Data:           []float64{2.5, 3.5, 4.5},
		},
	}

	for _, tC := range testcases {
		a := sum(tC.Data)

		if a != tC.ExpectedResult {
			t.Errorf("expected %f, got %f", tC.ExpectedResult, a)
		}
	}
}

func TestAvg(t *testing.T) {
	testcases := []struct {
		ExpectedResult float64
		Data           []float64
	}{
		{
			ExpectedResult: 0,
			Data:           []float64{},
		},
		{
			ExpectedResult: 2,
			Data:           []float64{1, 2, 3},
		},
		{
			ExpectedResult: -1,
			Data:           []float64{-2, 0},
		},
		{
			ExpectedResult: 2.5,
			Data:           []float64{2, 3},
		},
	}

	for _, tC := range testcases {
		a := avg(tC.Data)

		if a != tC.ExpectedResult {
			t.Errorf("expected %f, got %f", tC.ExpectedResult, a)
		}
	}
}
