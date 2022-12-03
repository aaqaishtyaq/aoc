package day01

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
)

//go:embed example_input input
var fs embed.FS

func readInput(filename string) ([]int, error) {
	data, err := utils.ReadInputFile(fs)
	if err != nil {
		return nil, err
	}

	inputText := string(data)
	inputSplit := strings.Split(inputText, "\n")
	nums := make([]int, len(inputSplit))
	for i, inputString := range inputSplit {
		if inputString == "" {
			break
		}
		nums[i], err = strconv.Atoi(inputString)
		if err != nil {
			return nil, err
		}
	}

	return nums, err
}

func Solution() error {
	input, err := readInput(utils.InputFile)
	if err != nil {
		return err
	}

	start := time.Now()
	part1Result := part1(input, 1)
	part2Result := part2(input, 3)
	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)
	return nil
}

func part1(nums []int, windowSize int) int {
	return solver(nums, windowSize)
}

func part2(nums []int, windowSize int) int {
	return solver(nums, windowSize)
}

func solver(nums []int, windowSize int) int {
	increment := 0

	i, j := 0, windowSize
	for j < len(nums) {
		if nums[j] > nums[i] {
			increment++
		}
		i++
		j++
	}

	return increment
}
