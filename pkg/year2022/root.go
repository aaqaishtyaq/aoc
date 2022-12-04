package year2022

import (
	"github.com/aaqaishtyaq/aoc/pkg/registry"
	"github.com/aaqaishtyaq/aoc/pkg/year2022/day01"
	"github.com/aaqaishtyaq/aoc/pkg/year2022/day02"
	"github.com/aaqaishtyaq/aoc/pkg/year2022/day03"
	"github.com/aaqaishtyaq/aoc/pkg/year2022/day04"
)

func Register() registry.Registry {
	registery := registry.NewRegistry()

	registery.Add(1, day01.Solution)
	registery.Add(2, day02.Solution)
	registery.Add(3, day03.Solution)
	registery.Add(4, day04.Solution)

	return registery
}
