package utils

import (
	"errors"
	"os"
)

func ParseInputFileName() string {
	args := os.Args[2:]
	if len(args) != 1 {
		panic(errors.New("expected an input file name"))
	}

	return args[0]
}
