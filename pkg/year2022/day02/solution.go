package day02

import (
	"embed"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aaqaishtyaq/aoc/pkg/utils"
)

//go:embed example_input input
var fs embed.FS

var (
	Rock    = "Rock"
	Paper   = "Paper"
	Scissor = "Scissor"
	Win     = "Win"
	Draw    = "Draw"
	Lose    = "Lose"
)

var (
	gameScore = map[string]int{
		Rock:    1,
		Paper:   2,
		Scissor: 3,
	}

	handToOutcome = map[string]string{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	handToStep = map[string]string{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
	}

	gameWinPts = map[string]string{
		Rock:    Scissor,
		Paper:   Rock,
		Scissor: Paper,
	}

	gameLoosePts = map[string]string{
		Rock:    Paper,
		Paper:   Scissor,
		Scissor: Rock,
	}

	outcomeScore = map[string]int{
		Win:  6,
		Draw: 3,
		Lose: 0,
	}
)

func Solution() error {
	rounds, err := readFile()
	if err != nil {
		return err
	}

	start := time.Now()
	part1Result, part2Result := 0, 0

	for _, r := range rounds {
		turns := strings.Split(r, " ")
		elf1, elf2 := turns[0], turns[1]

		s1, err := part01(elf1, elf2)
		if err != nil {
			return err
		}

		s2, err := part02(elf1, elf2)
		if err != nil {
			return err
		}

		part1Result += s1
		part2Result += s2
	}

	end := time.Since(start)

	fmt.Println("Solution for Part 1 : ", part1Result)
	fmt.Println("Solution for Part 2 : ", part2Result)
	fmt.Println("Time                : ", end)

	return nil
}

func part02(turn1, turn2 string) (score int, err error) {
	firstPlayerHand := handToStep[turn1]
	secondplayerHand := ""
	expectedOutcome := handToOutcome[turn2]

	switch expectedOutcome {
	case Draw:
		secondplayerHand = firstPlayerHand
	case Lose:
		secondplayerHand = gameWinPts[firstPlayerHand]
	case Win:
		secondplayerHand = gameLoosePts[firstPlayerHand]
	default:
	}

	score += outcomeScore[expectedOutcome]
	score += gameScore[secondplayerHand]

	return score, nil
}

func part01(turn1, turn2 string) (score int, err error) {
	handToStep := map[string]string{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}

	step1, ok := handToStep[turn1]
	if !ok {
		err = errors.New("unable to find step for a hand")

		return
	}

	step2, ok := handToStep[turn2]
	if !ok {
		err = errors.New("unable to find step for a hand")

		return
	}

	// add initial hand score
	score = score + gameScore[step2]

	// add outcome score
	if step1 == step2 {
		score += outcomeScore[Draw]

		return
	}

	if gameLoosePts[step1] == step2 {
		score += outcomeScore[Win]
	} else {
		score += outcomeScore[Lose]
	}

	return
}

func readFile() ([]string, error) {
	data, err := utils.ReadInputFile(fs)
	if err != nil {
		return nil, err
	}

	inputText := string(data)
	rounds := strings.Split(inputText, "\n")
	return rounds, nil
}
