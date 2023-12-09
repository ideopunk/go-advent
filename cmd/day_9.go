package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ideopunk/advent/convert"
)

func NextValue(pattern []int) int {
	allZero := true
	for _, entry := range pattern {
		if entry != 0 {
			allZero = false
			break
		}
	}

	if allZero {
		return 0
	}

	newPattern := make([]int, len(pattern)-1)

	for i := 1; i < len(pattern); i++ {
		newPattern[i-1] = pattern[i] - pattern[i-1]
	}

	nextValue := NextValue(newPattern)
	return pattern[len(pattern)-1] + nextValue
}

// forgive self for duplicating 
func PrevValue(pattern []int) int {
	allZero := true
	for _, entry := range pattern {
		if entry != 0 {
			allZero = false
			break
		}
	}

	if allZero {
		return 0
	}

	newPattern := make([]int, len(pattern)-1)

	for i := 1; i < len(pattern); i++ {
		newPattern[i-1] = pattern[i] - pattern[i-1]
	}

	prevValue := PrevValue(newPattern)
	return pattern[0] - prevValue
}

func Day9(lines *[]string, pt int) (string, error) {

	sum := 0

	if pt == 1 {
		for _, line := range *lines {
			pattern, err := convert.StringSliceToIntSlice(strings.Split(line, " "))
			if err != nil {
				return "", fmt.Errorf("could not convert line to pattern: %v", err)
			}

			sum += NextValue(pattern)
		}
	} else if pt == 2 {
		for _, line := range *lines {
			pattern, err := convert.StringSliceToIntSlice(strings.Split(line, " "))
			if err != nil {
				return "", fmt.Errorf("could not convert line to pattern: %v", err)
			}

			sum += PrevValue(pattern)
		}
	}

	return strconv.Itoa(sum), nil
}
