package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/aaqaishtyaq/aoc/days/day01"
	_ "github.com/aaqaishtyaq/aoc/days/day02"
	"github.com/aaqaishtyaq/aoc/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a day number.")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Unable to parse day number.")
		os.Exit(1)
	}

	f, ok := registry.Registry[day]
	if !ok {
		fmt.Printf("Not implemented for day: %d\n", day)
		os.Exit(1)
	}

	f()
}
