package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id      int
	red     int
	green   int
	blue    int
	product int
}

func Day2(part int) (string, error) {
	lines, err := fileToArr("./inputs/day_2.txt")
	if err != nil {
		return "", fmt.Errorf("could not convert file to arr: %v", err)
	}

	r, err := regexp.Compile(", |; ")
	if err != nil {
		return "", fmt.Errorf("could not compile regex: %v", err)
	}

	games := make([]Game, 0)

	for _, line := range lines {

		// get ID
		split := strings.Split(line, ": ")
		n := strings.Replace(split[0], "Game ", "", -1)
		id, err := strconv.Atoi(n)

		if err != nil {
			return "", fmt.Errorf("could not convert id to int: %v", err)
		}

		game := Game{
			id: id,
		}

		// get colors
		colorPhrases := r.Split(split[1], -1)

		for _, colorPhrase := range colorPhrases {
			colorSplit := strings.Split(colorPhrase, " ")
			count := colorSplit[0]
			color := colorSplit[1]

			switch color {
			case "red":
				red, err := strconv.Atoi(count)
				if err != nil {
					return "", fmt.Errorf("could not convert red to int: %v", err)
				}

				if red > game.red {
					game.red = red
				}
			case "green":
				green, err := strconv.Atoi(count)
				if err != nil {
					return "", fmt.Errorf("could not convert green to int: %v", err)
				}

				if green > game.green {
					game.green = green
				}
			case "blue":
				blue, err := strconv.Atoi(count)
				if err != nil {
					return "", fmt.Errorf("could not convert blue to int: %v", err)
				}

				if blue > game.blue {
					game.blue = blue
				}
			default:
				return "", fmt.Errorf("unknown color: %v", color)
			}
		}

		// Part One!
		// decide if valid
		if part == 1 {
			if game.red > 12 || game.green > 13 || game.blue > 14 {
				continue
			}
		}

		// Part Two!
		// calculate product
		game.product = game.red * game.green * game.blue

		games = append(games, game)
	}

	sum := 0

	for _, game := range games {
		if part == 1 {
			// part one
			sum += game.id
		} else {
			// part two
			sum += game.product
		}
	}

	return strconv.Itoa(sum), nil
}

func Day2PartTwo() (string, error) {
	return "", nil
}
