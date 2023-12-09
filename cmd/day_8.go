package cmd

import (
	"strconv"
	"strings"
)

type Node struct {
	name string
	l    string
	r    string
}

func Day8(lines *[]string) (string, error) {
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
}
