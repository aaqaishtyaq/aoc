package day02

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
	registry.Registry[2] = main
}

type Line struct {
	direction string
	distance  int
}

func readInput(filename string) []Line {
	filePath, err := filepath.Abs("../days/day02/" + filename)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}


	inputText := string(data)
	inputLines := strings.Split(inputText, "\n")

	lines := make([]Line, len(inputLines))
	for i, inputLine := range inputLines {
		lineSplit := strings.Split(inputLine, " ")
		distance, _ := strconv.Atoi(lineSplit[1])
		lines[i] = Line{
			direction: lineSplit[0],
			distance:  distance,
		}
	}
	return lines
}

func main() {
	filename := utils.ParseInputFileName()
	input := readInput(filename)

	start := time.Now()
	part1Result := part1(input)
	part2Result := part2(input)
	end := time.Since(start)


	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)
}

func part1(lines []Line) int {
	x, y := 0, 0
	for _, line := range lines {
		switch line.direction {
		case "forward":
			x += line.distance
		case "down":
			y -= line.distance
		case "up":
			y += line.distance
		}
	}

	depth := y * -1
	return x * depth
}

func part2(lines []Line) int {
	x, y, aim := 0, 0, 0

	for _, line := range lines {
		switch line.direction {
		case "forward":
			x += line.distance
			y -= aim * line.distance
		case "down":
			aim += line.distance
		case "up":
			aim -= line.distance
		}
	}

	depth := y * -1
	return x * depth
}
