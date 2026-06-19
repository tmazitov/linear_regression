package parsing

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

type CSVFile struct {
	records	[][2]int
}

func NewCSVFile(filePath string) (*CSVFile, error) {

	records, err := parseCSVFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("CSVFile error: %w", err)
	}

	return &CSVFile{
		records: records,
	}, nil
}

func (c *CSVFile) Records() [][2]int {
	return c.records
}

func parseCSVFile(filePath string) ([][2]int, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("parseCSVFile error: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parseCSVFile error: %w", err)
	}
	
	var records [][2]int
	for _, line := range lines[1:] {
		record, err := lineToRecord(line)
		if err != nil {
			return nil, fmt.Errorf("parseCSVFile error: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
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