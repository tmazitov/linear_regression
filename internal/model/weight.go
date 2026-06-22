package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Weight struct {
	K        float64 `json:"k"`
	B        float64 `json:"b"`
	DistMin  float64 `json:"dist_min"`
	DistMax  float64 `json:"dist_max"`
	PriceMin float64 `json:"price_min"`
	PriceMax float64 `json:"price_max"`
}

func NewWeight() *Weight {
	return &Weight{}
}

func (w *Weight) Update(k, b, distMin, distMax, priceMin, priceMax float64) {
	w.K = k
	w.B = b
	w.DistMin = distMin
	w.DistMax = distMax
	w.PriceMin = priceMin
	w.PriceMax = priceMax
}

func (w *Weight) LinearCoefficients() (k, b float64) {
	k = w.K * (w.PriceMax - w.PriceMin) / (w.DistMax - w.DistMin)
	b = w.B*(w.PriceMax-w.PriceMin) + w.PriceMin - k*w.DistMin
	return
}

func (w *Weight) Save(filePath string) error {
	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return fmt.Errorf("weight save: %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("weight save: %w", err)
	}
	return nil
}
