package cmd

import (
	"fmt"
	"strconv"
)

func Day1() (string, error) {
	lines, err := fileToArr("./inputs/day_1.txt")
	if err != nil {
		return "", fmt.Errorf("could not convert file to arr: %v", err)
	}

	sum := 0
	for _, line := range lines {
		var char1 string
		var char2 string

		for _, char := range line {
			runeToInt, err := strconv.Atoi(string(char))

			if err != nil {
				continue
			}

			char1 = strconv.Itoa(runeToInt)
			break
		}

		for i := len(line) - 1; i >= 0; i = i - 1 {
			runeToInt, err := strconv.Atoi(string(line[i]))

			if err != nil {
				continue
			}

			char2 = strconv.Itoa(runeToInt)
			break
		}

		combined, err := strconv.Atoi(char1 + char2)

		if err != nil {
			return "", fmt.Errorf("could not combine chars 1 and 2: %v", err)
		}

		sum += combined
	}

	return strconv.Itoa(sum), nil
}
