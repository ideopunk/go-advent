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

	switch os.Args[1] {
	case "1":
		results, err := cmd.Day1()
		if err != nil {
			fmt.Println("could not run day 1: ", err)
			return
		}
		fmt.Println(results)
	default:
		fmt.Println("no day provided")
	}

}
