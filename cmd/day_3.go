package cmd

import (
	"fmt"
	"regexp"
	"strconv"
)

// 140x140

type Index struct {
	y      int
	x      int
	number int
}

func getFullNumber(line string, yIndex, xIndex int) (Index, error) {
	// get the number
	num := string(line[xIndex])

	// add afterward
	for i := xIndex + 1; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			break
		}

		num += string(line[i])
	}

	fullNumber := Index{y: yIndex}

	// add before
	for i := xIndex - 1; i >= 0; i-- {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			fullNumber.x = i + 1
			break
		}

		num = string(line[i]) + num
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return Index{}, fmt.Errorf("could not convert final numstring %v to int: %v", num, err)
	}

	fullNumber.number = n

	return fullNumber, nil
}

func Day3() (string, error) {
	lines, err := fileToArr("./inputs/day_3.txt")
	if err != nil {
		return "", fmt.Errorf("could not convert file to arr: %v", err)
	}

	// find all special symbols
	areaIndices := []Index{}

	r, err := regexp.Compile("[^\\w\\.\\d ]")
	if err != nil {
		return "", fmt.Errorf("could not compile regex: %v", err)
	}

	for i, line := range lines {
		lineIndices := r.FindAllStringIndex(line, -1)

		for _, index := range lineIndices {
			// add their area to the list

			// above
			areaIndices = append(areaIndices, Index{y: i - 1, x: index[0] - 1})
			areaIndices = append(areaIndices, Index{y: i - 1, x: index[0]})
			areaIndices = append(areaIndices, Index{y: i - 1, x: index[0] + 1})

			// beside
			areaIndices = append(areaIndices, Index{y: i, x: index[0] - 1})
			areaIndices = append(areaIndices, Index{y: i, x: index[0] + 1})

			// below
			areaIndices = append(areaIndices, Index{y: i + 1, x: index[0] - 1})
			areaIndices = append(areaIndices, Index{y: i + 1, x: index[0]})
			areaIndices = append(areaIndices, Index{y: i + 1, x: index[0] + 1})
		}
	}

	// find any numbers in their area.
	fullNumbers := map[string]Index{}

	for _, index := range areaIndices {
		// is this a number?
		_, err := strconv.Atoi(string(lines[index.y][index.x]))
		if err != nil {
			continue
		}

		// if this is a number, get the extension of the number and its starting index
		num, err := getFullNumber(lines[index.y], index.y, index.x)
		if err != nil {
			return "", fmt.Errorf("could not get full number: %v", err)
		}

		sx := strconv.Itoa(num.x)
		sy := strconv.Itoa(num.y)
		fullNumbers[sx+"-"+sy] = num
	}

	// sum
	sum := 0
	for _, num := range fullNumbers {
		sum += num.number
	}

	return strconv.Itoa(sum), nil
}

func Day3Part2() (string, error) {
	lines, err := fileToArr("./inputs/day_3.txt")
	if err != nil {
		return "", fmt.Errorf("could not convert file to arr: %v", err)
	}

	// find all * symbols

	r, err := regexp.Compile("\\*")
	if err != nil {
		return "", fmt.Errorf("could not compile regex: %v", err)
	}

	sum := 0
	for i, line := range lines {
		gearIndices := r.FindAllStringIndex(line, -1)
		for _, index := range gearIndices {
			partIndices := map[string]Index{}
			// add their area to the list

			XY := []Index{
				// above
				{y: i - 1, x: index[0] - 1},
				{y: i - 1, x: index[0]},
				{y: i - 1, x: index[0] + 1},
				//beside
				{y: i, x: index[0] - 1},
				{y: i, x: index[0] + 1},
				// below
				{y: i + 1, x: index[0] - 1},
				{y: i + 1, x: index[0]},
				{y: i + 1, x: index[0] + 1},
			}

			// get all the full numbers in the area.
			for _, xy := range XY {
				// is this a number?
				_, err := strconv.Atoi(string(lines[xy.y][xy.x]))
				if err != nil {
					continue
				}

				fullNum, err := getFullNumber(lines[xy.y], xy.y, xy.x)
				if err != nil {
					return "", fmt.Errorf("could not get full number: %v", err)
				}

				partIndices[strconv.Itoa(fullNum.x)+"-"+strconv.Itoa(fullNum.y)] = fullNum

			}

			// only allow gears with two parts c:
			if len(partIndices) != 2 {
				continue
			}

			product := 1
			for _, num := range partIndices {
				product *= num.number
			}

			sum += product
		}

	}

	return strconv.Itoa(sum), nil
}
