package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"lucascron.com/advent-of-code-2023/util"
)

const (
	redMax   = 12
	greenMax = 13
	blueMax  = 14

	gamePrefix = "Game"

	blueSuffix  = "blue"
	greenSuffix = "green"
	redSuffix   = "red"
)

var (
	files = [2]string{
		"example.txt",
		"input.txt",
	}
	colors = [3]string{blueSuffix, greenSuffix, redSuffix}
)

type Roll struct {
	reds   int
	greens int
	blues  int
}

func (r *Roll) setRollFromSuffix(s string, value int) {
	switch s {
	case blueSuffix:
		r.blues = value
	case redSuffix:
		r.reds = value
	case greenSuffix:
		r.greens = value
	}
}

func (r *Roll) printRoll() {
	fmt.Println(fmt.Sprintf("Reds: %d, Blues: %d, Greens: %d", r.reds, r.blues, r.greens))
}

func (r *Roll) parseValues(values []string) {
	for _, value := range values {
		for _, c := range colors {
			if strings.Contains(value, c) {
				value = strings.ReplaceAll(value, c, "")
				valueInt, err := strconv.Atoi(value)
				util.Check(err)
				r.setRollFromSuffix(c, valueInt)
			}
		}
	}
	r.printRoll()
}

type Game struct {
	rolls    []Roll
	id       int
	MinRed   int
	MinBlue  int
	MinGreen int
}

func (g *Game) isPossible() bool {
	isPossible := true
	for _, roll := range g.rolls {
		isPossible = isPossible && roll.reds <= redMax && roll.blues <= blueMax && roll.greens <= greenMax
		if !isPossible {
			break
		}
	}
	return isPossible
}

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)

		var idSum int
		var minProductSum int

		for _, line := range lines {
			// fmt.Println(line)
			spacelessLine := strings.ReplaceAll(line, " ", "")
			gameAndOutput := strings.Split(spacelessLine, ":")

			if len(gameAndOutput) > 2 {
				log.Fatal("Something is busted")
			}

			gameId := gameAndOutput[0][len(gamePrefix):]
			gameIdInt, err := strconv.Atoi(gameId)
			util.Check(err)

			game := Game{
				id: gameIdInt,
			}
			rolls := strings.Split(gameAndOutput[1], ";")
			for _, roll := range rolls {
				var r Roll
				values := strings.Split(roll, ",")
				r.parseValues(values)

				if r.blues != 0 && game.MinBlue < r.blues {
					game.MinBlue = r.blues
				}
				if r.greens != 0 && game.MinGreen < r.greens {
					game.MinGreen = r.greens
				}
				if r.reds != 0 && game.MinRed < r.reds {
					game.MinRed = r.reds
				}

				game.rolls = append(game.rolls, r)
			}

			minProduct := game.MinBlue * game.MinGreen * game.MinRed
			minProductSum += minProduct

			isPossible := game.isPossible()
			if isPossible {
				idSum += gameIdInt
			}
			fmt.Println(fmt.Sprintf("Game ID: %d, Is Possible: %t, Current ID Sum: %d, Min Red: %d, Min Blue: %d, Min Green: %d, Min Product: %d",
				gameIdInt, isPossible, idSum, game.MinRed, game.MinBlue, game.MinGreen, minProduct))
		}

		fmt.Println(fmt.Sprintf("Path: %s, ID Sum: %d, Min Product Sum: %d", path, idSum, minProductSum))
	}
}
