package day03

import (
	"embed"
	"fmt"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
)

//go:embed *input
var fs embed.FS

type set struct {
	in [53]bool
}

func (s *set) add(a []byte) {
	for _, c := range a {
		if c == 13 {
			continue
		}

		s.in[getIndex(c)] = true
	}
}

func (s *set) exist(c byte) bool {
	if c == 13 {
		return false
	}

	return s.in[getIndex(c)]
}

func getIndex(c byte) byte {
	var i byte = 0
	if c >= 'a' && c <= 'z' {
		i = c - 'a'
	} else {
		i = c - 'A' + 26
	}

	return i + 1
}

func Solution() error {
	compartments, err := readFile()
	if err != nil {
		return nil
	}

	start := time.Now()
	part1Result := part01(compartments)
	part2Result := part02(compartments)
	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)

	return nil
}

func part01(cts []string) int32 {
	var sum int32

	for _, c := range cts {
		s := &set{}
		l := len(c)
		s.add([]byte(c[:l/2]))

		for _, b := range c[l/2:] {
			if s.exist(byte(b)) {
				sum += int32(getIndex(byte(b)))

				break
			}
		}
	}

	return sum
}

func part02(cts []string) int32 {
	var sum int32

	for i := 0; i < len(cts)/3; i++ {
		s1, s2, s3 := &set{}, &set{}, &set{}
		s1.add([]byte(cts[i*3]))
		s2.add([]byte(cts[i*3+1]))
		s3.add([]byte(cts[i*3+2]))

		for j := 1; j < 53; j++ {
			if s1.in[j] && s2.in[j] && s3.in[j] {
				sum += int32(j)
			}
		}
	}

	return sum
}

func readFile() ([]string, error) {
	data, err := utils.ReadInputFile(fs)
	if err != nil {
		return nil, err
	}

	ipText := strings.TrimSpace(string(data))
	ipItems := strings.Split(ipText, "\n")

	return ipItems, nil
}
