package util

import (
	"bufio"
	"embed"
)

func GetStrings(fileB embed.FS, fileName string) ([]string, error) {
	var lines []string

	file, err := fileB.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
