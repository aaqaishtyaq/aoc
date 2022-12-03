package cmd

import (
	"fmt"
	"strconv"

	"github.com/aaqaishtyaq/aoc/pkg/year2021"
	"github.com/spf13/cobra"
)

func new2021Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2021",
		Short: "Advent of Code 2021 solutions",
		Args:  cobra.MinimumNArgs(1),
		RunE:  Run2021,
	}

	return cmd
}

func Run2021(cmd *cobra.Command, args []string) error {
	day, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("unable to parse day number: %v", err)
	}

	registery := year2021.Register()

	f, ok := registery.Get(day)
	if !ok {
		return fmt.Errorf("not implemented for day: %d", day)
	}

	if err := f(); err != nil {
		return err
	}

	return nil
}
