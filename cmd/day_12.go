package cmd

import (
	"fmt"
	"strings"

	"github.com/ideopunk/advent/convert"
)

type Record struct {
	original     string
	arrangements map[string]bool
	sizes        []int
}

func (r *Record) Groups() []string {
	groups := []string{}

	for _, r := range r.original {
		s := string(r)
		if s == "?" || s == "#" {
			groups[len(groups)-1] += s
		} else {
			// start a new group
			if groups[len(groups)-1] != "" {
				groups = append(groups, "")
			}
		}
	}

	return groups
}

func NewRecord(s string) (Record, error) {
	p := strings.Split(s, " ")

	preInts := strings.Split(p[1], ",")
	ints, err := convert.StringSliceToIntSlice(preInts)
	if err != nil {
		return Record{}, fmt.Errorf("failed to convert string slice to int slice: %w", err)
	}

	r := Record{
		original: p[0],
		sizes:    ints,
	}

	return r, nil
}

func (r *Record) Match() {
	groups := r.Groups()
	sizes := r.sizes
	fmt.Println(groups, sizes)

	// for _, group := range groups {
		
	// }
}

func Day12(lines *[]string) (string, error) {

	return "", nil
}

func Day12Part2(lines *[]string) (string, error) {
	return "", nil
}
