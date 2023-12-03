package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func FileToArr(filename string) ([]string, error) {
	if filename == "" {
		return nil, fmt.Errorf("no filename provided")
	}

	file, err := os.Open(filename) // For read access.
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}
