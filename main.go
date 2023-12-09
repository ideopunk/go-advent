package main

import (
	"fmt"
	"os"

	"github.com/ideopunk/advent/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no day provided")
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("no part provided")
		return
	}

	var results string
	var err error

	lines, err := cmd.FileToArr("./inputs/day_" + os.Args[1] + ".txt")
	if err != nil {
		fmt.Printf("could not convert file to arr: %v", err)
		return
	}

	switch os.Args[1] {
	case "1":
		if os.Args[2] == "1" {
			results, err = cmd.Day1(&lines)
		} else {
			results, err = cmd.Day1Part2(&lines)
		}
	case "2":
		if os.Args[2] == "1" {
			results, err = cmd.Day2(&lines, 1)
		} else {
			results, err = cmd.Day2(&lines, 2)
		}
	case "3":
		if os.Args[2] == "1" {
			results, err = cmd.Day3(&lines)
		} else {
			results, err = cmd.Day3Part2(&lines)
		}
	case "4":
		if os.Args[2] == "1" {
			results, err = cmd.Day4(&lines)
		} else {
			results, err = cmd.Day4Part2(&lines)
		}
	case "5":
		if os.Args[2] == "1" {
			results, err = cmd.Day5(&lines, 1)
		} else {
			results, err = cmd.Day5(&lines, 2)
		}
	case "6":
		if os.Args[2] == "1" {
			results, err = cmd.Day6(&lines)
		} else {
			results, err = cmd.Day6Part2(&lines)
		}
	case "7":
		if os.Args[2] == "1" {
			results, err = cmd.Day7(&lines, 1)
		} else {
			results, err = cmd.Day7(&lines, 2)
		}
	case "8":
		if os.Args[2] == "1" {
			results, err = cmd.Day8(&lines, 1)
		} else {
			results, err = cmd.Day8(&lines, 2)
		}
	default:
		fmt.Println("no matching day found")
	}

	if err != nil {
		fmt.Printf("just couldn't handle day %v part %v: %v", os.Args[1], os.Args[2], err)
		return
	}

	fmt.Println(results)
}
