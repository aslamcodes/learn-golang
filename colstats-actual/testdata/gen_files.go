package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	numFiles    = 5000
	rowsPerFile = 2500
	outDir      = "testdata/benchmark"
)

var areas = []string{
	"chennai", "delhi", "mumbai", "kolkata", "bengaluru",
	"hyderabad", "pune", "kochi", "trivandrum",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := os.MkdirAll(outDir, 0755); err != nil {
		panic(err)
	}

	for i := 1; i <= numFiles; i++ {
		path := filepath.Join(outDir, fmt.Sprintf("sample_%d.csv", i))
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		w := csv.NewWriter(f)
		_ = w.Write([]string{"temperature", "area"})

		for j := 0; j < rowsPerFile; j++ {
			temp := -10 + rand.Float64()*50 // -10 to 50
			area := areas[rand.Intn(len(areas))]
			_ = w.Write([]string{
				strconv.FormatFloat(temp, 'f', 2, 64),
				area,
			})
		}

		w.Flush()
		f.Close()
	}

	fmt.Println("Done")
}
