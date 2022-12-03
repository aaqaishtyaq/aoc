package utils

import "embed"

var InputFile string

func ReadInputFile(fs embed.FS) ([]byte, error) {
	return fs.ReadFile(InputFile)
}
