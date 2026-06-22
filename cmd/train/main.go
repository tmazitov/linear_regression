package main

import (
	"fmt"
	"os"

	"github.com/tmazitov/linear_regression/internal/model"
	"github.com/tmazitov/linear_regression/internal/parsing"
	"github.com/tmazitov/linear_regression/internal/plot"
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

	csvFile.Normalize()
	prices := csvFile.NormPrices()
	distances := csvFile.NormDistances()

	m := model.NewModel()

	m.Train(model.Dataset{
		Prices:    prices.Values(),
		Distances: distances.Values(),
		PriceMin:  prices.Min(),
		PriceMax:  prices.Max(),
		DistMin:   distances.Min(),
		DistMax:   distances.Max(),
	}, 0.5, 1000)

	m.Weight().Save("./weights.json")

	originalPrices := csvFile.Prices()
	originalDistances := csvFile.Distances()

	points := []*plot.Point{}
	for index, price := range originalPrices {
		points = append(points, plot.NewPoint(originalDistances[index], price))
	}

	p, err := plot.NewPlot(points)
	if err != nil {
		fatal(err)
	}
	k, b := m.Weight().LinearCoefficients()
	p.AddLine(k, b)
	p.Show()
}
