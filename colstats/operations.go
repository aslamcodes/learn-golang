package main

import "fmt"

type StatsFunc func(data []float64) float64

func getOperation(op string) (StatsFunc, error) {
	switch op {
	case "sum":
		return sum, nil
	case "avg":
		return avg, nil
	}
	return nil, fmt.Errorf("Invalid operation id")
}

func sum(data []float64) float64 {
	var totalSum float64
	for _, val := range data {
		totalSum += val
	}
	return totalSum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}
