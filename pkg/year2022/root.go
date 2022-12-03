package year2022

import (
	"github.com/aaqaishtyaq/aoc/pkg/year2022/day01"
	"github.com/aaqaishtyaq/aoc/pkg/registry"
)

func Register() registry.Registry {
	registery := registry.NewRegistry()

	registery.Add(1, day01.Solution)

	return registery
}
