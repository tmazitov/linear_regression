package main

import (
	"fmt"
	"os"

	"github.com/tmazitov/linear_regression/internal/model"
)

func fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func main() {
	var distance float64
	fmt.Print("Enter mileage: ")
	if _, err := fmt.Scan(&distance); err != nil {
		fatal(fmt.Errorf("invalid mileage value: %w", err))
	}

	m := model.NewModel()
	if err := m.LoadWeights("./weights.json"); err != nil {
		fatal(err)
	}

	fmt.Printf("Estimated price for %.0f km: %.2f\n", distance, m.EstimatePrice(distance))
}
