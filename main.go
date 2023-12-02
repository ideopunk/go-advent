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

	switch os.Args[1] {
	case "1":
		if os.Args[2] == "1" {
			results, err = cmd.Day1()
		} else {
			results, err = cmd.Day1PartTwo()
		}
	case "2":
		if os.Args[2] == "1" {
			results, err = cmd.Day2(1)
		} else {
			results, err = cmd.Day2(2)
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
