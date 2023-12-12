package cmd

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X    int
	Y    int
	char string
}

func NextPipe(pipes *[]string, prev, curr Point) (Point, Point, error) {
	currPipe := (*pipes)[curr.Y][curr.X]
	newCurr := Point{0, 0, ""}

	switch currPipe {
	case '|':
		if prev.Y == curr.Y+1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1, string((*pipes)[curr.Y-1][curr.X])}
		} else if prev.Y == curr.Y-1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1, string((*pipes)[curr.Y+1][curr.X])}
		}
	case 'J':
		if prev.X == curr.X-1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1, string((*pipes)[curr.Y-1][curr.X])}
		} else if prev.Y == curr.Y-1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y, string((*pipes)[curr.Y][curr.X-1])}
		}
	case 'L':
		if prev.X == curr.X+1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1, string((*pipes)[curr.Y-1][curr.X])}
		} else if prev.Y == curr.Y-1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y, string((*pipes)[curr.Y][curr.X+1])}
		}
	case 'F':
		if prev.X == curr.X+1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1, string((*pipes)[curr.Y+1][curr.X])}
		} else if prev.Y == curr.Y+1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y, string((*pipes)[curr.Y][curr.X+1])}
		}
	case '7':
		if prev.X == curr.X-1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1, string((*pipes)[curr.Y+1][curr.X])}
		} else if prev.Y == curr.Y+1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y, string((*pipes)[curr.Y][curr.X-1])}
		}
	case '-':

		if prev.X == curr.X+1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y, string((*pipes)[curr.Y][curr.X-1])}
		} else if prev.X == curr.X-1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y, string((*pipes)[curr.Y][curr.X+1])}
		}
	case 'S':
		return prev, curr, fmt.Errorf("well, how did I get here?")
	}

	if newCurr.X == 0 && newCurr.Y == 0 {
		return curr, newCurr, fmt.Errorf("could not find next pipe, prev: %v, curr: %v", string((*pipes)[prev.Y][prev.X]), string((*pipes)[curr.Y][curr.X]))
	}
	return curr, newCurr, nil
}

// note that we are modifying as we go, smoothing down the "*" characters that have been added along the way.
// Between our turners we want pure "-" and "|" characters, in order to form accurate walls for the inner dots.
func NextPipeLong(pipes *[]string, prev, curr Point) ([]Point, error) {
	turners := []string{"J", "L", "F", "7", "S"}

	newPipes := []Point{}
	c := curr

	xDir := 0
	yDir := 0
	switch (*pipes)[curr.Y][curr.X] {
	case '|':
		if prev.Y == curr.Y+1 {
			// we're going up
			yDir = -1
		} else if prev.Y == curr.Y-1 {
			// we're going down
			yDir = 1
		}
	case 'J':
		if prev.X == curr.X-1 {
			// we're going up
			yDir = -1
		} else if prev.Y == curr.Y-1 {
			// we're going left
			xDir = -1
		}
	case 'L':
		if prev.X == curr.X+1 {
			// we're going up
			yDir = -1
		} else if prev.Y == curr.Y-1 {
			// we're going right
			xDir = 1
		}
	case 'F':
		if prev.X == curr.X+1 {
			// we're going down
			yDir = 1
		} else if prev.Y == curr.Y+1 {
			// we're going right
			xDir = 1
		}
	case '7':
		if prev.X == curr.X-1 {
			// we're going down
			yDir = 1
		} else if prev.Y == curr.Y+1 {
			// we're going left
			xDir = -1
		}
	case '-':
		if prev.X == curr.X+1 {
			// we're going left
			xDir = -1
		} else if prev.X == curr.X-1 {
			// we're going right
			xDir = 1
		}
	case 'S':
		// this should already be handled outside this function
		return newPipes, fmt.Errorf("well, how did I get here?")
	}

	for {
		// get next one, add it to our great list.
		c = Point{c.X + xDir, c.Y + yDir, string((*pipes)[c.Y+yDir][c.X+xDir])}
		newPipes = append(newPipes, c)

		// and if we've hit a turner, it's a wrap
		if slices.Contains(turners, string((*pipes)[c.Y][c.X])) {
			break
		}

		// if this isn't a turner, then we need to modify it (turn asterisk into line)
		if yDir != 0 {
			// mutate
			(*pipes)[c.Y] = (*pipes)[c.Y][:c.X] + "|" + (*pipes)[c.Y][c.X+1:]
		} else if xDir != 0 {
			// mutate
			(*pipes)[c.Y] = (*pipes)[c.Y][:c.X] + "-" + (*pipes)[c.Y][c.X+1:]
		}

	}

	if xDir == 0 && yDir == 0 {
		return newPipes, fmt.Errorf("could not find direction for next pipe, prev: %v, curr: %v", string((*pipes)[prev.Y][prev.X]), string((*pipes)[curr.Y][curr.X]))
	}

	return newPipes, nil
}

func IncreaseGridResolution(grid *[]string) []string {
	reAll := regexp.MustCompile(`(.)`)

	expandedGrid := make([]string, len(*grid)*2)
	// expand the grid horizintally
	for i := range *grid {
		s := reAll.ReplaceAllString((*grid)[i], "${1}*")
		expandedGrid[i*2] = s
	}

	// expand the grid vertically
	for i := 0; i < len(*grid); i++ {
		expandedGrid[i*2+1] = reAll.ReplaceAllString(expandedGrid[0], "*")
	}

	return expandedGrid
}

func FindStarter(pipes *[]string) (Point, error) {
	s := Point{0, 0, "S"}
	for i, line := range *pipes {
		for j, char := range line {
			if char == 'S' {
				s = Point{j, i, "S"}
			}
		}
	}

	if s.X == 0 && s.Y == 0 {
		return s, fmt.Errorf("could not find starter")
	}

	return s, nil
}

func FindFirstConnector(lines *[]string, prev Point) Point {
	curr := Point{0, 0, string((*lines)[prev.Y][prev.X])}

	above := (*lines)[prev.Y-1][prev.X]
	below := (*lines)[prev.Y+1][prev.X]
	left := (*lines)[prev.Y][prev.X-1]
	right := (*lines)[prev.Y][prev.X+1]

	switch above {
	case '|':
		curr = Point{prev.X, prev.Y - 1, ""}
	case 'F':
		curr = Point{prev.X, prev.Y - 1, ""}
	case '7':
		curr = Point{prev.X, prev.Y - 1, ""}
	}

	switch below {
	case '|':
		curr = Point{prev.X, prev.Y + 1, ""}
	case 'J':
		curr = Point{prev.X, prev.Y + 1, ""}
	case 'L':
		curr = Point{prev.X, prev.Y + 1, ""}
	}

	switch left {
	case '-':
		curr = Point{prev.X - 1, prev.Y, ""}
	case 'F':
		curr = Point{prev.X - 1, prev.Y, ""}
	case 'L':
		curr = Point{prev.X - 1, prev.Y, ""}
	}

	switch right {
	case '-':
		curr = Point{prev.X + 1, prev.Y, ""}
	case 'J':
		curr = Point{prev.X + 1, prev.Y, ""}
	case '7':
		curr = Point{prev.X + 1, prev.Y, ""}
	}

	curr.char = string((*lines)[curr.Y][curr.X])
	return curr
}

// and mutate along the way
func FindFirstConnectorLong(lines *[]string, prev Point) ([]Point, error) {
	above := (*lines)[prev.Y-2][prev.X]
	below := (*lines)[prev.Y+2][prev.X]
	left := (*lines)[prev.Y][prev.X-2]
	right := (*lines)[prev.Y][prev.X+2]

	switch above {
	case '|', 'F', '7':
		(*lines)[prev.Y-1] = (*lines)[prev.Y-1][:prev.X] + "|" + (*lines)[prev.Y-1][prev.X+1:]
		return []Point{{prev.X, prev.Y - 1, ""}, {prev.X, prev.Y - 2, ""}}, nil
	}

	switch below {
	case '|', 'J', 'L':
		(*lines)[prev.Y+1] = (*lines)[prev.Y+1][:prev.X] + "|" + (*lines)[prev.Y+1][prev.X+1:]
		return []Point{{prev.X, prev.Y + 1, ""}, {prev.X, prev.Y + 2, ""}}, nil
	}

	switch left {
	case '-', 'F', 'L':
		(*lines)[prev.Y] = (*lines)[prev.Y][:prev.X-1] + "-" + (*lines)[prev.Y][prev.X:]
		return []Point{{prev.X - 1, prev.Y, ""}, {prev.X - 2, prev.Y, ""}}, nil
	}

	switch right {
	case '-', 'J', '7':
		(*lines)[prev.Y] = (*lines)[prev.Y][:prev.X+1] + "-" + (*lines)[prev.Y][prev.X+2:]
		return []Point{{prev.X + 1, prev.Y, ""}, {prev.X + 2, prev.Y, ""}}, nil
	}

	return []Point{}, fmt.Errorf("could not find first connector")
}

func Day10(lines *[]string) (string, error) {
	// find S
	prev, err := FindStarter(lines)
	if err != nil {
		return "", err
	}

	s := prev

	// find one of its two connectors
	curr := FindFirstConnector(lines, prev)

	// then get going
	count := 1
	for {
		count++

		// just in case
		if count > 20000 {
			break
		}

		prev, curr, err = NextPipe(lines, prev, curr)

		if err != nil {
			return "", fmt.Errorf("could not find next pipe: %v", err)
		}
		// we've wrapped around
		if curr.X == s.X && curr.Y == s.Y {
			break
		}
	}

	// might not need to account for odd numbers?
	return strconv.Itoa(count / 2), nil
}

func Day10Part2(lines *[]string) (string, error) {

	// first, mark outside
	smallPrev, err := FindStarter(lines)
	smallS := smallPrev
	if err != nil {
		return "", fmt.Errorf("could not find first starter: %v", err)
	}

	smallCurr := FindFirstConnector(lines, smallPrev)

	count := 0

	smallPipes := []Point{smallPrev, smallCurr}
	for {
		count++

		// just in case
		if count > 20000 {
			return "", fmt.Errorf("count got too high: %v", count)
		}

		smallPrev, smallCurr, err = NextPipe(lines, smallPrev, smallCurr)

		// let's track our pipes eh?

		if err != nil {
			return "", fmt.Errorf("could not find next pipe: %v", err)
		}

		smallPipes = append(smallPipes, smallCurr)

		// we've wrapped around
		if smallPipes[len(smallPipes)-1].X == smallS.X && smallPipes[len(smallPipes)-1].Y == smallS.Y {
			break
		}
	}

	// mark all our pipes
	for _, pipe := range smallPipes {
		line := (*lines)[pipe.Y]
		newLine := line[:pipe.X] + "o" + line[pipe.X+1:]
		(*lines)[pipe.Y] = newLine
	}

	// mark non-pipes
	re := regexp.MustCompile(`[^o]`)

	for i, line := range *lines {
		(*lines)[i] = re.ReplaceAllString(line, ".")
	}

	// return pipes to old values
	for _, pipe := range smallPipes {
		line := (*lines)[pipe.Y]
		newLine := line[:pipe.X] + pipe.char + line[pipe.X+1:]
		(*lines)[pipe.Y] = newLine
	}

	// now work with higher resolution!
	bigGrid := IncreaseGridResolution(lines)

	// find S again
	prev, err := FindStarter(&bigGrid)
	if err != nil {
		return "", fmt.Errorf("could not find second starter: %v", err)
	}

	s := prev

	// find one of its two connectors

	// manual labor
	bigGrid[prev.Y+1] = bigGrid[prev.Y+1][:prev.X] + "|" + bigGrid[prev.Y+1][prev.X+1:]

	pipes := []Point{{prev.X, prev.Y + 1, ""}, {prev.X, prev.Y + 2, ""}}

	// then get going filling out the pipes and grabbing all their new locations
	count = 0
	for {
		count++

		// just in case
		if count > 20000 {
			return "", fmt.Errorf("count got too high: %v", count)
		}

		newPipes, err := NextPipeLong(&bigGrid, pipes[len(pipes)-2], pipes[len(pipes)-1])

		// let's track our pipes eh?

		if err != nil {
			return "", fmt.Errorf("could not find next pipe: %v", err)
		}

		pipes = append(pipes, newPipes...)

		// we've wrapped around
		if pipes[len(pipes)-1].X == s.X && pipes[len(pipes)-1].Y == s.Y {
			break
		}
	}

	// clean the outside
	outerDots := map[Point]bool{}
	outerDots[Point{0, 0, ""}] = false

	nonPipes := []string{".", "*"}
	// flood
	for {
		// get the ones that haven't been scanned
		toCheck := []Point{}
		for p, dot := range outerDots {
			if !dot {
				toCheck = append(toCheck, p)
			}
		}

		for _, p := range toCheck {
			// add new folks
			if p.Y != 0 {
				// check above
				_, ok := outerDots[Point{p.X, p.Y - 1, ""}]
				if !ok {
					above := bigGrid[p.Y-1][p.X]
					if slices.Contains(nonPipes, string(above)) {
						outerDots[Point{p.X, p.Y - 1, ""}] = false
					}
				}
			}

			if p.Y != len(bigGrid)-1 {
				// check below
				_, ok := outerDots[Point{p.X, p.Y + 1, ""}]
				if !ok {
					below := bigGrid[p.Y+1][p.X]
					if slices.Contains(nonPipes, string(below)) {
						outerDots[Point{p.X, p.Y + 1, ""}] = false
					}
				}
			}

			if p.X != 0 {
				// check left
				_, ok := outerDots[Point{p.X - 1, p.Y, ""}]
				if !ok {
					left := bigGrid[p.Y][p.X-1]
					if slices.Contains(nonPipes, string(left)) {
						outerDots[Point{p.X - 1, p.Y, ""}] = false
					}
				}
			}

			if p.X != len((bigGrid)[p.Y])-1 {
				// check right
				_, ok := outerDots[Point{p.X + 1, p.Y, ""}]
				if !ok {
					right := bigGrid[p.Y][p.X+1]
					if slices.Contains(nonPipes, string(right)) {
						outerDots[Point{p.X + 1, p.Y, ""}] = false
					}
				}
			}

			// checked
			outerDots[p] = true
		}

		// we got em all
		if len(toCheck) == 0 {
			break
		}
	}

	// replace outer dots with space
	for dot := range outerDots {
		bigGrid[dot.Y] = bigGrid[dot.Y][:dot.X] + " " + bigGrid[dot.Y][dot.X+1:]
	}

	// now check remaining dots
	countRemaining := 0
	for i := range bigGrid {
		countRemaining += strings.Count(bigGrid[i], ".")
	}

	return strconv.Itoa(countRemaining), nil
}
