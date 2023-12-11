package cmd

import (
	"fmt"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func NextPipe(pipes *[]string, prev, curr Point) (Point, Point, error) {
	currPipe := (*pipes)[curr.Y][curr.X]
	newCurr := Point{0, 0}

	switch currPipe {
	case '|':
		if prev.Y == curr.Y+1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1}
		} else if prev.Y == curr.Y-1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1}
		}
	case 'J':
		if prev.X == curr.X-1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1}
		} else if prev.Y == curr.Y-1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y}
		}
	case 'L':
		if prev.X == curr.X+1 {
			// we're going up
			newCurr = Point{curr.X, curr.Y - 1}
		} else if prev.Y == curr.Y-1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y}
		}
	case 'F':
		if prev.X == curr.X+1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1}
		} else if prev.Y == curr.Y+1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y}
		}
	case '7':
		if prev.X == curr.X-1 {
			// we're going down
			newCurr = Point{curr.X, curr.Y + 1}
		} else if prev.Y == curr.Y+1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y}
		}
	case '-':

		if prev.X == curr.X+1 {
			// we're going left
			newCurr = Point{curr.X - 1, curr.Y}
		} else if prev.X == curr.X-1 {
			// we're going right
			newCurr = Point{curr.X + 1, curr.Y}
		}
	case 'S':
		return prev, curr, fmt.Errorf("well, how did I get here?")
	}

	if newCurr.X == 0 && newCurr.Y == 0 {
		return curr, newCurr, fmt.Errorf("could not find next pipe, prev: %v, curr: %v", string((*pipes)[prev.Y][prev.X]), string((*pipes)[curr.Y][curr.X]))
	}
	return curr, newCurr, nil
}

func FindStarter(pipes *[]string) Point {
	s := Point{0, 0}
	for i, line := range *pipes {
		for j, char := range line {
			if char == 'S' {
				s = Point{j, i}
			}
		}
	}
	return s
}

func FindFirstConnector(lines *[]string, prev Point) Point {
	curr := Point{0, 0}

	above := (*lines)[prev.Y-1][prev.X]
	below := (*lines)[prev.Y+1][prev.X]
	left := (*lines)[prev.Y][prev.X-1]
	right := (*lines)[prev.Y][prev.X+1]

	switch above {
	case '|':
		curr = Point{prev.X, prev.Y - 1}
	case 'F':
		curr = Point{prev.X, prev.Y - 1}
	case '7':
		curr = Point{prev.X, prev.Y - 1}
	}

	switch below {
	case '|':
		curr = Point{prev.X, prev.Y + 1}
	case 'J':
		curr = Point{prev.X, prev.Y + 1}
	case 'L':
		curr = Point{prev.X, prev.Y + 1}
	}

	switch left {
	case '-':
		curr = Point{prev.X - 1, prev.Y}
	case 'F':
		curr = Point{prev.X - 1, prev.Y}
	case 'L':
		curr = Point{prev.X - 1, prev.Y}
	}

	switch right {
	case '-':
		curr = Point{prev.X + 1, prev.Y}
	case 'J':
		curr = Point{prev.X + 1, prev.Y}
	case '7':
		curr = Point{prev.X + 1, prev.Y}
	}

	return curr
}

func Day10(lines *[]string) (string, error) {
	// find S
	prev := FindStarter(lines)
	s := prev

	// find one of its two connectors
	curr := FindFirstConnector(lines, prev)

	// then get going
	count := 1
	err := error(nil)
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
	return "", fmt.Errorf("not implemented")
}
