package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) Count() int {
	lilMap := map[rune]int{}

	for _, c := range h.cards {
		lilMap[c]++
	}

	return len(lilMap)
}

func (h *Hand) highestCount() int {
	lilMap := map[rune]int{}

	for _, c := range h.cards {
		lilMap[c]++
	}

	max := 0
	for _, v := range lilMap {
		if v > max {
			max = v
		}
	}

	return max
}

func faceToNum(face string) int {
	switch face {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		// we're just gonna ignore the possibility of this here fellow messing up :)
		num, err := strconv.Atoi(face)
		if err != nil {
			panic(fmt.Errorf(`Don't break my converter!`))
		}
		
		return num
	}
}

func SortHands(handA, handB Hand) int {
	for i := 0; i < len(handA.cards); i++ {
		preA := string(handA.cards[i])
		a := faceToNum(preA)

		preB := string(handB.cards[i])
		b := faceToNum(preB)
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}

	return 0
}

func Day7(lines *[]string) (string, error) {
	fives := []Hand{}
	fours := []Hand{}
	fullHouses := []Hand{}
	threesOfAKind := []Hand{}
	twoPairs := []Hand{}
	onePairs := []Hand{}
	highCards := []Hand{}

	// categorize lines by hands
	for _, line := range *lines {
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			return "", fmt.Errorf("could not convert bid to int: %v", err)
		}

		hand := Hand{cards: split[0], bid: bid}

		switch hand.Count() {
		case 1:
			fives = append(fives, hand)
		case 2:
			// distinguish ex AA8AA from 23332
			maximum := hand.highestCount()
			if maximum == 4 {
				fours = append(fours, hand)
			} else {
				fullHouses = append(fullHouses, hand)
			}
		case 3:
			// distinguish ex TTT98 from 998TT
			maximum := hand.highestCount()
			if maximum == 3 {
				threesOfAKind = append(threesOfAKind, hand)
			} else {
				twoPairs = append(twoPairs, hand)
			}
		case 4:
			onePairs = append(onePairs, hand)
		default:
			highCards = append(highCards, hand)
		}

	}
	// order hands internally
	slices.SortFunc(fives, SortHands)
	slices.SortFunc(fours, SortHands)
	slices.SortFunc(fullHouses, SortHands)
	slices.SortFunc(threesOfAKind, SortHands)
	slices.SortFunc(twoPairs, SortHands)
	slices.SortFunc(onePairs, SortHands)
	slices.SortFunc(highCards, SortHands)

	// now order them all by rank
	allHands := append([]Hand{}, highCards...)
	allHands = append(allHands, onePairs...)
	allHands = append(allHands, twoPairs...)
	allHands = append(allHands, threesOfAKind...)
	allHands = append(allHands, fullHouses...)
	allHands = append(allHands, fours...)
	allHands = append(allHands, fives...)

	// multiply bid, and sum
	sum := 0
	for i, hand := range allHands {
		winnings := hand.bid * (i + 1)
		sum += winnings
	}

	return strconv.Itoa(sum), nil
}
