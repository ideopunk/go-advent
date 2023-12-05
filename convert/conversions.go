package convert

import (
	"fmt"
	"strconv"
)

func StringSliceToIntSlice(s []string) ([]int, error) {

	ints := make([]int, len(s))
	for i, str := range s {
		integer, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("could not convert string to int: %v", err)
		}
		ints[i] = integer
	}
	return ints, nil
}
