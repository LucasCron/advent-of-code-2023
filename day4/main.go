package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"lucascron.com/advent-of-code-2023/util"
)

var (
	files = [2]string{
		"example.txt",
		"input.txt",
	}
)

type Game struct {
	id              int
	winners         []string
	numbersReceived []string
	winCount        int
	points          float64
	copyIdsWon      []int
	cardsFromCopies int
}

func (g *Game) parseWinCount() {
	for _, w := range g.numbersReceived {
		if slices.Contains(g.winners, w) {
			g.winCount += 1
		}

		if g.winCount > 0 {
			g.points = math.Pow(2, float64(g.winCount-1))
		}
	}
}

func main() {
	for _, path := range files {
		lines := util.ReadTextFileToArray(path)
		games := map[int]Game{}
		var pointSum float64

		for _, line := range lines {
			gameAndResults := strings.Split(line, ":")
			gameNumber, err := strconv.Atoi(strings.ReplaceAll(gameAndResults[0], " ", "")[len("Card"):])
			util.Check(err)

			results := gameAndResults[1]
			winnersAndReceived := strings.Split(results, "|")
			winners := strings.Split(strings.ReplaceAll(strings.TrimSpace(winnersAndReceived[0]), "  ", " "), " ")
			received := strings.Split(strings.ReplaceAll(strings.TrimSpace(winnersAndReceived[1]), "  ", " "), " ")
			game := Game{
				id:              gameNumber,
				winners:         winners,
				numbersReceived: received,
			}
			game.parseWinCount()
			games[gameNumber] = game
			// fmt.Println(fmt.Sprintf("Game: %d, Wins: %d, Points: %g", game.id, game.winCount, game.points))
			pointSum += game.points
		}
		fmt.Println(fmt.Sprintf("Total Points: %g", pointSum))

		totalCards := len(games)
		for i := len(games); i > 0; i-- {
			game := games[i]

			for j := game.winCount; j > 0; j-- {
				game.copyIdsWon = append(game.copyIdsWon, i+j)
				game.cardsFromCopies += 1
				game.cardsFromCopies += games[i+j].cardsFromCopies
			}

			games[i] = game
			totalCards += game.cardsFromCopies
		}
		fmt.Println(fmt.Sprintf("Total Cards: %d", totalCards))
	}
}
