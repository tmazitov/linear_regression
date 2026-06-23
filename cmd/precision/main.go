package main

import (
	"fmt"
	"os"

	"github.com/tmazitov/linear_regression/internal/model"
	"github.com/tmazitov/linear_regression/internal/parsing"
)

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func main() {
	csvFilePath, err := parsing.GetCSVFilePath()
	if err != nil {
		fatal(err)
	}

	csvFile, err := parsing.NewCSVFile(csvFilePath)
	if err != nil {
		fatal(err)
	}

	m := model.NewModel()
	if err := m.LoadWeights("./weights.json"); err != nil {
		fatal(err)
	}

	prices := csvFile.Prices()
	distances := csvFile.Distances()
	n := len(prices)

	var precision float64
	for i := range n {
		predicted := m.EstimatePrice(distances[i])
		diff := prices[i] - predicted
		if diff < 0 {
			diff = -diff
		}
		precision += diff / prices[i]
	}
	precision = precision / float64(n) * 100

	fmt.Printf("Precision: %.2f%%\n", 100-precision)
}
