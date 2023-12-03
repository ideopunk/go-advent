package cmd

import (
	"fmt"
	"regexp"
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

func wordToNum(word string) (string, error) {
	switch word {
	case "one":
		return "1", nil
	case "two":
		return "2", nil
	case "three":
		return "3", nil
	case "four":
		return "4", nil
	case "five":
		return "5", nil
	case "six":
		return "6", nil
	case "seven":
		return "7", nil
	case "eight":
		return "8", nil
	case "nine":
		return "9", nil
	}

	return "0", fmt.Errorf("could not convert word to num: %v", word)
}

func Day1Part2() (string, error) {
	lines, err := fileToArr("./inputs/day_1.txt")
	if err != nil {
		return "", fmt.Errorf("could not convert file to arr: %v", err)
	}

	r, err := regexp.Compile("one|two|three|four|five|six|seven|eight|nine|[1-9]")
	if err != nil {
		return "", fmt.Errorf("could not compile regex: %v", err)
	}

	sum := 0
	for _, line := range lines {
		first := r.FindString(line)
		if len(first) > 1 {
			var err error
			first, err = wordToNum(first)
			if err != nil {
				return "", fmt.Errorf("could not convert first match to int: %v", err)
			}
		}

		var last string

		length := len(line)
		for i := len(line) - 1; i >= 0; i = i - 1 {
			last = r.FindString(line[i:length])
			if len(last) > 0 {
				break
			}
		}

		if len(last) > 1 {
			var err error
			last, err = wordToNum(last)
			if err != nil {
				return "", fmt.Errorf("could not convert first match to int: %v", err)
			}
		}

		combined, err := strconv.Atoi(first + last)
		if err != nil {
			return "", fmt.Errorf("could not combine : %v", err)
		}

		sum += combined
	}

	return strconv.Itoa(sum), nil
}
