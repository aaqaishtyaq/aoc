package cmd

import (
	"context"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
	"github.com/spf13/cobra"
)

func ExecuteContext(ctx context.Context) error {
	rootCmd := newRootCommand()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}

func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "aoc",
		Short: "Advent of Code in Goland",
	}
	rootCmd.PersistentFlags().StringVarP(&utils.InputFile, "file", "f", "example_input", "input filename")

	rootCmd.AddCommand(new2021Command())
	rootCmd.AddCommand(new2022Command())

	return rootCmd
}
