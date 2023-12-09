package cmd

import (
	"math"
	"strconv"
	"strings"
)

type Node struct {
	name string
	l    string
	r    string
}

func PrimeFactor(n int) map[int]int {
	factors := []int{}

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	m := map[int]int{}

	for _, i := range factors {
		_, present := m[i]
		if !present {
			m[i] = 1
		} else {
			m[i]++
		}
	}
	return m
}

func PrimeFactorCombo(primeArr []map[int]int) int {

	highestPowerFactors := map[int]int{}

	for _, primes := range primeArr {
		for i, exp := range primes {

			_, present := highestPowerFactors[i]
			if !present {
				highestPowerFactors[i] = exp
				continue
			}

			if exp > highestPowerFactors[i] {
				highestPowerFactors[i] = exp
			}
		}
	}

	product := 1
	for b, exp := range highestPowerFactors {
		product *= int(math.Pow(float64(b), float64(exp)))
	}
	return product
}

func Day8(lines *[]string, pt int) (string, error) {
	count := 0

	nodeMap := make(map[string]Node)

	// map out these nodes
	for i := 2; i < len(*lines); i++ {
		line := strings.Replace((*lines)[i], ")", "", -1)
		parts := strings.Split(line, " = (")
		name := parts[0]
		lr := strings.Split(parts[1], ", ")

		n := Node{
			name: name,
			l:    lr[0],
			r:    lr[1],
		}

		nodeMap[name] = n
	}

	currNode := nodeMap["AAA"]

	// keep wrapping if we haven't gotten it yet
	commands := strings.Split((*lines)[0], "")
	i := 0

	if pt == 1 {
		for i < len(commands) {
			count++

			if commands[i] == "L" {
				currNode = nodeMap[currNode.l]
			} else {
				currNode = nodeMap[currNode.r]
			}

			i++
			if i == len(commands) {
				i = 0
			}

			if currNode.name == "ZZZ" {
				break
			}
		}

		return strconv.Itoa(count), nil

	} else {
		// part 2
		counts := [6]int{0, 0, 0, 0, 0, 0}

		currNodes := []Node{}

		// find all the ones that end in A
		for _, n := range nodeMap {
			if n.name[2] == 'A' {
				currNodes = append(currNodes, n)
			}
		}

		for i < len(commands) {
			count++

			// move it
			for j := 0; j < len(currNodes); j++ {
				if commands[i] == "L" {
					currNodes[j] = nodeMap[currNodes[j].l]
				} else {
					currNodes[j] = nodeMap[currNodes[j].r]
				}
			}

			// increment it
			i++
			if i == len(commands) {
				i = 0
			}

			// get counts for each so we can LCM
			for j, n := range currNodes {
				if n.name[2] == 'Z' && counts[j] == 0 {
					counts[j] = count
				}
			}

			anyNotZ := false

			for _, n := range counts {
				if n == 0 {
					anyNotZ = true
					break
				}
			}

			// we have all of them, now find lowest common multiple
			if !anyNotZ {
				primes := []map[int]int{}
				for _, count = range counts {
					primes = append(primes, PrimeFactor(count))
				}

				combined := PrimeFactorCombo(primes)
				return strconv.Itoa(combined), nil
			}

		}

		return "", nil
	}

}
