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

func (m *Map) LineToRun(s string) error {
	split := strings.Split(s, " ")
	ints, err := convert.StringSliceToIntSlice(split)
	if err != nil {
		return fmt.Errorf("could not convert lineToRun strs into ints: %w", err)
	}

	m.runs = append(m.runs, Run{ints[0], ints[1], ints[2]})
	m.Sort()
	return nil
}

func (m *Map) Sort() {
	sort.Slice(m.runs, func(i, j int) bool {
		return m.runs[i].destination < m.runs[j].destination
	})
}

// assumes already sorted
func (m *Map) Trickle(seed int) int {

	n := seed
	// is it within one of our runs? map over.
	for _, run := range m.runs {
		if seed >= run.source && seed <= run.source+run.len {
			diff := seed - run.source
			n = run.destination + diff
			break
		}
	}

	return n
}

func Day5(lines *[]string) (string, error) {
	// seeds
	preseeds := (*lines)[0]
	firstSplit := strings.Split(preseeds, ": ")
	seedStrings := strings.Split(firstSplit[1], " ")
	seeds, err := convert.StringSliceToIntSlice(seedStrings)

	if err != nil {
		return "", fmt.Errorf("could not convert seed strings to ints: %w", err)
	}

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
		err := maps[len(maps)-1].LineToRun(line)

		if err != nil {
			return "", fmt.Errorf("could not run line: %w", err)
		}
	}

	var lowest int
	for _, seed := range seeds {
		curr := seed
		for _, m := range maps {
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

	return strconv.Itoa(lowest), nil
}

func Day5Part2(lines *[]string) (string, error) {
	return "", nil
}
