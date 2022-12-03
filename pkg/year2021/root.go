package year2021

import (
	"github.com/aaqaishtyaq/aoc/pkg/year2021/day01"
	"github.com/aaqaishtyaq/aoc/pkg/year2021/day02"
	"github.com/aaqaishtyaq/aoc/pkg/registry"
)

func Register() registry.Registry {
	registery := registry.NewRegistry()

	registery.Add(1, day01.Solution)
	registery.Add(2, day02.Solution)

	return registery
}
