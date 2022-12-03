package day01

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
)

//go:embed example_input input
var fs embed.FS

func Solution() error {
	cals, err := readInput()
	if err != nil {
		return err
	}

	start := time.Now()
	part1Result := part01(cals)
	part2Result := part02(cals)
	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)

	return nil
}

func part01(input []int) int {
	type maxNum struct{ idx, size int }

	var max = maxNum{}
	for i, cal := range input {
		if i == 0 {
			max = maxNum{i, cal}
		} else {
			if cal > max.size {
				max = maxNum{i, cal}
			}
		}
	}

	return max.size
}

func part02(input []int) int {
	sort.Slice(input, func(i, j int) bool { return input[i] > input[j] })
	return input[0] + input[1] + input[2]
}

func readInput() ([]int, error) {
	data, err := utils.ReadInputFile(fs)
	if err != nil {
		return nil, err
	}

	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")
	var totalCals = []int{}

	total := 0
	lastIdx := 0
	for _, l := range inputLines {
		if l == "" {
			totalCals = append(totalCals, total)
			lastIdx += 1
			total = 0
		} else {
			cal, err := strconv.Atoi(l)
			if err != nil {
				return nil, fmt.Errorf("unable to convert string to int: %s", l)
			}
			total += cal
		}
	}

	return totalCals, nil
}
