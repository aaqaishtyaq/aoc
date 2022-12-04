package day04

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
)

//go:embed *input
var fs embed.FS

type sectionAssignment struct {
	start int
	end   int
}

func Solution() error {
	items, err := readFile()
	if err != nil {
		return nil
	}

	start := time.Now()
	part1Result := part01(items)
	part2Result := part02(items)
	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)

	return nil
}

func part01(items [][]sectionAssignment) int {
	res := 0
	for _, pair := range items {
		if pair[0].start <= pair[1].start && pair[0].end >= pair[1].end {
			res++
		} else if pair[0].start >= pair[1].start && pair[0].end <= pair[1].end {
			res++
		}
	}

	return res
}

func part02(items [][]sectionAssignment) int {
	res := 0
	for _, pair := range items {
		if pair[1].start >= pair[0].start && pair[1].start <= pair[0].end {
			res++
		} else if pair[0].start >= pair[1].start && pair[0].start <= pair[1].end {
			res++
		}
	}

	return res
}

func readFile() ([][]sectionAssignment, error) {
	data, err := utils.ReadInputFile(fs)
	if err != nil {
		return nil, err
	}

	ipText := strings.TrimSpace(string(data))
	pairs := strings.Split(ipText, "\n")

	items := [][]sectionAssignment{}

	for _, p := range pairs {
		pairSplit := strings.Split(p, ",")
		assignments := []sectionAssignment{}

		for _, i := range pairSplit {
			e := strings.Split(i, "-")
			first, err := strconv.Atoi(e[0])
			if err != nil {
				return nil, err
			}

			second, err := strconv.Atoi(e[1])
			if err != nil {
				return nil, err
			}

			assignments = append(assignments, sectionAssignment{
				start: first,
				end:   second,
			})
		}

		items = append(items, assignments)
	}

	return items, nil
}
