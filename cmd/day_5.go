package cmd

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/ideopunk/advent/convert"
)

type Run struct {
	destination int
	source      int
	len         int
}
type Map struct {
	runs []Run
}

func (m *Map) LineToRun(s string, sortStyle string) error {
	split := strings.Split(s, " ")
	ints, err := convert.StringSliceToIntSlice(split)
	if err != nil {
		return fmt.Errorf("could not convert lineToRun strs into ints: %w", err)
	}

	m.runs = append(m.runs, Run{ints[0], ints[1], ints[2]})

	if sortStyle == "source" {
		m.SortSource()
	} else {
		m.SortDestination()
	}
	return nil
}

func (m *Map) SortSource() {
	sort.Slice(m.runs, func(i, j int) bool {
		return m.runs[i].source < m.runs[j].source
	})
}

func (m *Map) SortDestination() {
	sort.Slice(m.runs, func(i, j int) bool {
		return m.runs[i].destination < m.runs[j].destination
	})
}

// assumes already sorted
func (m *Map) Trickle(seed int) int {

	n := seed

	// is it after all our runs? map over.
	highest := m.runs[len(m.runs)-1].source + m.runs[len(m.runs)-1].len
	if seed > highest {
		return n
	}

	for _, run := range m.runs {
		// we skipped over it, we're done!
		if seed < run.source {
			return n
		}

		// is it within one of our runs? nice! over.
		if seed >= run.source && seed < run.source+run.len {
			diff := run.destination - run.source
			n = seed + diff
			break
		}
	}
	return n
}

func (m *Map) TrickleUp(num int) int {
	for _, run := range m.runs {
		if num >= run.destination && num < run.destination+run.len {
			diff := num - run.destination
			return run.source + diff
		}
	}
	return num
}

func (m *Map) HighestPossibleValue() int {
	hpv := 0
	for _, run := range m.runs {
		highestRunValue := run.destination + run.len
		if highestRunValue > hpv {
			hpv = highestRunValue
		}
	}
	return hpv
}

func Day5(lines *[]string, pt int) (string, error) {
	// seeds
	seeds := []int{}

	preseeds := (*lines)[0]
	firstSplit := strings.Split(preseeds, ": ")
	seedStrings := strings.Split(firstSplit[1], " ")
	s, err := convert.StringSliceToIntSlice(seedStrings)
	if err != nil {
		return "", fmt.Errorf("could not convert seed strings to ints: %w", err)
	}

	seeds = append(seeds, s...)

	// chunking
	maps := []Map{}

	re := regexp.MustCompile(`[a-zA-Z]`)
	for _, line := range (*lines)[1:] {
		//

		// start a new chunk if we're blank
		if len(line) == 0 {
			newMap := Map{}
			maps = append(maps, newMap)
			continue
		}

		// if line contains letters, continue
		if re.MatchString(line) {
			continue
		}

		// otherwise, use the runs to update the map
		style := "source"
		if pt == 2 {
			style = "destination"
		}

		err := maps[len(maps)-1].LineToRun(line, style)

		if err != nil {
			return "", fmt.Errorf("could not run line: %w", err)
		}
	}

	starterSeed := 0
	if pt == 1 {
		starterSeed = PartOneTrickle(&seeds, &maps)
	} else {
		starterSeed = PartTwoTrickle(&seeds, &maps)
	}

	return strconv.Itoa(starterSeed), nil
}

func PartOneTrickle(seeds *[]int, maps *[]Map) int {
	var lowest int

	for _, seed := range *seeds {
		curr := seed
		for _, m := range *maps {
			n := m.Trickle(curr)
			curr = n
		}

		if lowest == 0 {
			lowest = curr
			continue
		}

		if curr < lowest {
			lowest = curr
		}

	}

	return lowest
}

func PartTwoTrickle(seeds *[]int, maps *[]Map) int {
	seed := 0

	// get seeds in order
	seedPairs := [][]int{}
	for i := 0; i < len(*seeds)-1; i = i + 2 {
		seedPairs = append(seedPairs, []int{(*seeds)[i], (*seeds)[i+1]})
	}

	sort.Slice(seedPairs, func(i, j int) bool {
		return seedPairs[i][0] < seedPairs[j][0]
	})

	// move up through the waterfall
	bottomMap := (*maps)[len(*maps)-1]
	highestPossibleValue := bottomMap.HighestPossibleValue()

	for i := 0; i < highestPossibleValue; i++ {
		reverseVal := i
		for j := len(*maps) - 1; j >= 0; j-- {
			reverseVal = (*maps)[j].TrickleUp(reverseVal)
		}

		for _, pair := range seedPairs {
			if reverseVal >= pair[0] && reverseVal < pair[0]+pair[1] {
				seed = reverseVal
				return i
			}
		}
	}

	return seed
}
