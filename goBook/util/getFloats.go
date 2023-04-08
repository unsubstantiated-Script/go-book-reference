package util

import (
	"bufio"
	"embed"
	"strconv"
)

func GetFloats(fileB embed.FS, fileName string) ([]float64, error) {
	var numbers []float64

	file, err := fileB.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
