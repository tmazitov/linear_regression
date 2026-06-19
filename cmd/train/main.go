package main

import (
	"fmt"
	"os"
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

	csvFile.Normalize()
	prices := csvFile.NormPrices()
	distances := csvFile.NormDistances()
	fmt.Println("Normalized Prices:", prices)
	fmt.Println("Normalized Distances:", distances)
}
