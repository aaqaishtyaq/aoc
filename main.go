package main

import (
	"context"
	"os"

	"github.com/aaqaishtyaq/aoc/cmd"
	"golang.org/x/exp/slog"
)

func main() {
	lHandler := slog.NewTextHandler(os.Stdout)
	log := slog.New(lHandler)
	slog.SetDefault(log)

	if err := cmd.ExecuteContext(context.Background()); err != nil {
		log.Warn(err.Error())
		os.Exit(1)
	}
}
