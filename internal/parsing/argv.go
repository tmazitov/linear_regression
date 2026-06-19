package parsing

import (
	"os"
	"errors"
)

func GetCSVFilePath() (string, error) {

	argv := os.Args

	if len(argv) != 2 {
		return "", errors.New("GetCSVFilePath error: missing CSV file path argument")
	}

	return argv[1], nil
}
