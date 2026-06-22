package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tmazitov/linear_regression/internal/model"
)

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: predict <distance>")
		os.Exit(1)
	}

	distance, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fatal(fmt.Errorf("invalid distance value: %s", os.Args[1]))
	}

	m := model.NewModel()
	if err := m.LoadWeights("./weights.json"); err != nil {
		fatal(err)
	}

	fmt.Printf("Estimated price for %.0f km: %.2f\n", distance, m.EstimatePrice(distance))
}
