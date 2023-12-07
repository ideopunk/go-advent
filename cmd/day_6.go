package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func (r *Race) MaxSpeed() (int, bool) {
	even := true
	if r.time%2 != 0 {
		even = false
	}
	return r.time / 2, even
}

// square problem.
func (r *Race) WinningHoldTimes() int {
	count := 0

	speed, even := r.MaxSpeed()
	hold := speed
	if !even {
		hold++
	}

	// maybe we just aren't fast enough
	if int(speed*hold) < r.distance {
		return 0
	}

	// check against distance
	for {
		if int(speed*hold) <= r.distance {
			break
		}

		if int(speed*hold) > r.distance {
			count += 2
		}

		speed--
		hold++

	}

	// don't doublecount.
	if even {
		count--
	}

	return count
}

func LinesToRaces(lines *[]string) ([]Race, error) {
	races := []Race{}

	re, err := regexp.Compile("\\s+")
	if err != nil {
		return races, fmt.Errorf("could not compile regex: %v", err)
	}

	pretimes := strings.Split((*lines)[0], ":")[1]
	predistances := strings.Split((*lines)[1], ":")[1]

	times := re.Split(pretimes, -1)
	distances := re.Split(predistances, -1)

	// skip crud at beginning of regex
	for i := 1; i < len(times); i++ {
		t := times[i]
		d := distances[i]

		time, err := strconv.Atoi(t)
		if err != nil {
			return races, fmt.Errorf("could not convert time to int: %v", err)
		}
		distance, err := strconv.Atoi(d)
		if err != nil {
			return races, fmt.Errorf("could not convert distance to int: %v", err)
		}

		newRace := Race{time: time, distance: distance}
		races = append(races, newRace)
	}

	return races, nil
}

func SetToRace(lines *[]string) (Race, error) {
	race := Race{}
	re, err := regexp.Compile("\\s+")
	if err != nil {
		return race, fmt.Errorf("could not compile regex: %v", err)
	}

	pretime := strings.Split((*lines)[0], ":")[1]
	predistance := strings.Split((*lines)[1], ":")[1]

	t := re.ReplaceAllString(pretime, "")
	d := re.ReplaceAllString(predistance, "")

	time, err := strconv.Atoi(t)
	if err != nil {
		return race, fmt.Errorf("could not convert time to int: %v", err)
	}

	distance, err := strconv.Atoi(d)
	if err != nil {
		return race, fmt.Errorf("could not convert distance to int: %v", err)
	}

	race.time = time
	race.distance = distance

	return race, nil
}

func Day6(lines *[]string) (string, error) {
	races, err := LinesToRaces(lines)

	if err != nil {
		return "", fmt.Errorf("could not convert lines to races: %v", err)
	}

	product := 1
	for _, race := range races {
		holdTimes := race.WinningHoldTimes()
		product *= holdTimes
	}

	return strconv.Itoa(product), nil
}

func Day6Part2(lines *[]string) (string, error) {
	race, err := SetToRace(lines)

	if err != nil {
		return "", fmt.Errorf("could not convert set to races: %v", err)
	}

	holdTimes := race.WinningHoldTimes()

	return strconv.Itoa(holdTimes), nil
}
