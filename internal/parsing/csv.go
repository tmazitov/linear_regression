package parsing

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

type CSVFile struct {
	prices []int
	distances []int

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

func (c *CSVFile) Prices() 		[]int {return c.prices}
func (c *CSVFile) Distances()	[]int {return c.distances}
func (c *CSVFile) NormPrices() 	*NormRecords {return c.normPrices}
func (c *CSVFile) NormDistances() *NormRecords {return c.normDistances}

func parseCSVFile(filePath string) ([]int, []int, error) {

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
	
	var prices []int
	var distances []int
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

func lineToRecord(line []string) ([2]int, error) {
	if len(line) != 2 {
		return [2]int{}, fmt.Errorf("lineToRecord error: invalid line format")
	}
	distance, err := strconv.Atoi(line[0])
	if err != nil {
		return [2]int{}, fmt.Errorf("lineToRecord error: invalid distance value")
	}
	price, err := strconv.Atoi(line[1])
	if err != nil {
		return [2]int{}, fmt.Errorf("lineToRecord error: invalid price value")
	}
	return [2]int{distance, price}, nil
}

func (c *CSVFile) Normalize() {
	c.normPrices = NewNormRecords(c.prices)
	c.normDistances = NewNormRecords(c.distances)
}