package day01

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/registry"
	"github.com/aaqaishtyaq/aoc/utils"
)

func init() {
	registry.Registry[1] = main
}

func readInput(filename string) []int {
	filePath, err := filepath.Abs("../days/day01/" + filename)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputSplit := strings.Split(inputText, "\n")
	nums := make([]int, len(inputSplit))
	for i, inputString := range inputSplit {
		fmt.Printf("")
		nums[i], err = strconv.Atoi(inputString)
		if err != nil {
			panic(err)
		}
	}

	return nums
}

func main() {

	filename := utils.ParseInputFileName()
	input := readInput(filename)

	start := time.Now()
	part1Result := part1(input, 1)
	part2Result := part2(input, 3)
	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)
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
