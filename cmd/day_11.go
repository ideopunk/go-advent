package cmd

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Day11(lines *[]string) (string, error) {
	space := []string{}

	// Rows
	for i := 0; i < len(*lines); i++ {
		space = append(space, (*lines)[i])

		// add extra row if empty
		if !strings.Contains((*lines)[i], "#") {
			space = append(space, (*lines)[i])
		}
	}

	// Columns

	columnsToAdd := []int{}
	// find the columns we'll expand
	for i := 0; i < len((*lines)[0]); i++ {
		empty := true
		for j := 0; j < len(space); j++ {
			if space[j][i] == '#' {
				empty = false
				break
			}
		}

		if empty {
			columnsToAdd = append(columnsToAdd, i)
		}
	}

	// their indices will get messed up if we go upward, so we must reverse
	slices.Reverse(columnsToAdd)
	// insert columns
	for i := 0; i < len(space); i++ {
		for _, j := range columnsToAdd {
			space[i] = space[i][:j] + "." + space[i][j:]
		}
	}

	// find galaxies
	galaxies := []Coordinate{}

	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[i]); j++ {
			if space[i][j] == '#' {
				galaxies = append(galaxies, Coordinate{x: j, y: i})
			}
		}
	}

	// check and sum paths
	sum := 0

	// On^2 baby
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {

			// I think probably fine to not use abs, these are ordered.
			distance := math.Abs(float64(galaxies[j].x - galaxies[i].x))
			distance += math.Abs(float64(galaxies[j].y - galaxies[i].y))

			sum += int(distance)
		}
	}

	return strconv.Itoa(sum), nil
}

func Day11Part2(lines *[]string) (string, error) {
	space := *lines

	// Rows
	millionaireRows := []int{}

	// find the offending fellows
	for i := 0; i < len(space); i++ {
		if !strings.Contains(space[i], "#") {
			millionaireRows = append(millionaireRows, i)
		}
	}

	// Columns
	millionaireColumns := []int{}
	// find the columns we'll expand
	for i := 0; i < len(space[0]); i++ {
		empty := true
		for j := 0; j < len(space); j++ {
			if (space)[j][i] == '#' {
				empty = false
				break
			}
		}

		if empty {
			millionaireColumns = append(millionaireColumns, i)
		}
	}

	// find galaxies
	galaxies := []Coordinate{}

	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[i]); j++ {
			if space[i][j] == '#' {
				galaxies = append(galaxies, Coordinate{x: j, y: i})
			}
		}
	}

	// check and sum paths
	sum := 0

	// On^2 baby
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := 0
			gal1 := galaxies[i]
			gal2 := galaxies[j]

			// horizontal
			minH := min(gal1.x, gal2.x)
			maxH := max(gal1.x, gal2.x)

			for k := minH; k < maxH; k++ {
				// if it's one of the expanded columns, we're actually talking a milli
				if slices.Contains(millionaireColumns, k) {
					distance += 1_000_000
				} else {
					distance++
				}

			}

			// vertical
			minV := min(gal1.y, gal2.y)
			maxV := max(gal1.y, gal2.y)

			for k := minV; k < maxV; k++ {
				// if it's one of the expanded rows, we're actually talking a milli
				if slices.Contains(millionaireRows, k) {
					distance += 1_000_000
				} else {
					distance++
				}
			}

			sum += int(distance)
		}
	}

	return strconv.Itoa(sum), nil
}
