package parsing

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

type CSVFile struct {
	prices []float64
	distances []float64

	normPrices    *NormRecords
	normDistances *NormRecords
}

func NewCSVFile(filePath string) (*CSVFile, error) {

	prices, distances, err := parseCSVFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("CSVFile error: %w", err)
	}

	return &CSVFile{
		prices: prices,
		distances: distances,
	}, nil
}

func (c *CSVFile) Prices() 		[]float64 {return c.prices}
func (c *CSVFile) Distances()	[]float64 {return c.distances}
func (c *CSVFile) NormPrices() 	*NormRecords {return c.normPrices}
func (c *CSVFile) NormDistances() *NormRecords {return c.normDistances}

func parseCSVFile(filePath string) ([]float64, []float64, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("parseCSVFile error: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("parseCSVFile error: %w", err)
	}
	
	var prices []float64
	var distances []float64
	for _, line := range lines[1:] {
		record, err := lineToRecord(line)
		if err != nil {
			return nil, nil, fmt.Errorf("parseCSVFile error: %w", err)
		}
		prices = append(prices, record[1])
		distances = append(distances, record[0])
	}
	return prices, distances, nil
}

func lineToRecord(line []string) ([2]float64, error) {
	if len(line) != 2 {
		return [2]float64{}, fmt.Errorf("lineToRecord error: invalid line format")
	}
	distance, err := strconv.ParseFloat(line[0], 64)
	if err != nil {
		return [2]float64{}, fmt.Errorf("lineToRecord error: invalid distance value")
	}
	price, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return [2]float64{}, fmt.Errorf("lineToRecord error: invalid price value")
	}
	return [2]float64{distance, price}, nil
}

func (c *CSVFile) Normalize() {
	c.normPrices = NewNormRecords(c.prices)
	c.normDistances = NewNormRecords(c.distances)
}