package cmd

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ideopunk/advent/convert"
)

type Card struct {
	winners []int
	mine    []int
	matches int
}

func lineToCard(line string) (Card, error) {
	reExtra := regexp.MustCompile("Card.*?:")

	line = reExtra.ReplaceAllString(line, "")

	halves := strings.Split(line, "|")

	reSplit := regexp.MustCompile("\\s+")

	winnersStr := reSplit.Split(halves[0], -1)
	winners, err := convert.StringSliceToIntSlice(winnersStr[1 : len(winnersStr)-1])
	if err != nil {
		return Card{}, fmt.Errorf("could not convert winner strs to ints: %v", err)
	}

	mineStr := reSplit.Split(halves[1], -1)
	mine, err := convert.StringSliceToIntSlice(mineStr[1:])
	if err != nil {
		return Card{}, fmt.Errorf("could not convert mine strs to ints: %v", err)
	}

	return Card{
		winners: winners,
		mine:    mine,
		matches: 0,
	}, nil
}

// On^2 ;)
func (c *Card) play() {
	c.matches = 0
	for _, winner := range c.winners {
		for _, mine := range c.mine {
			if winner == mine {
				c.matches++
			}
		}
	}
}

func (c *Card) score() int {
	if c.matches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(c.matches-1)))
}

func Day4(lines *[]string) (string, error) {
	sum := 0
	for _, line := range *lines {
		card, err := lineToCard(line)
		if err != nil {
			return "", fmt.Errorf("could not convert line to card: %v", err)
		}
		card.play()
		score := card.score()
		sum += score
	}
	return strconv.Itoa(sum), nil
}

func Day4Part2(lines *[]string) (string, error) {

	return "", nil
}
