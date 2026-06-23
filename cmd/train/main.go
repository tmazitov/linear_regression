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
	
	csvFile.Normalize()
	prices := csvFile.NormPrices()
	distances := csvFile.NormDistances()

	m.SetupPlot(csvFile.Distances(), csvFile.Prices())
	m.Train(model.Dataset{
		Prices:    prices.Values(),
		Distances: distances.Values(),
		PriceMin:  prices.Min(),
		PriceMax:  prices.Max(),
		DistMin:   distances.Min(),
		DistMax:   distances.Max(),
	}, 0.5, 1000)
	m.UpdateLine()

	if err := m.ShowPlot(); err != nil {
		fatal(err)
	}

	if err := m.Weight().Save("./weights.json"); err != nil {
		fatal(err)
	}
}
