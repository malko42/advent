package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Round struct {
	red   int
	green int
	blue  int
}
type Game struct {
	id     int
	rounds []Round
}

type GameWithMinSet struct {
	*Game
	minSet Round
}

func newRound(red int, green int, blue int) Round {
	return Round{red, green, blue}
}

func newGame(id int, rounds []Round) Game {
	return Game{id, rounds}
}

func parseGame(line string) GameWithMinSet {
	var split = strings.Split(line, ":")
	var gameString = split[0]
	var roundsString = split[1]
	gameId, err := strconv.Atoi(strings.Split(gameString, " ")[1])
	check(err)
	var rounds = []Round{}
	var minSet = newRound(0, 0, 0)
	var splitRoundsString = strings.Split(roundsString, ";")
	for _, roundString := range splitRoundsString {
		gameRound := newRound(0, 0, 0)
		var splitRoundString = strings.Split(roundString, ",")
		for _, pioche := range splitRoundString {
			var splitPioche = strings.Split(pioche, " ")
			var color = splitPioche[2]
			var qty = splitPioche[1]
			switch color {
			case "red":
				var parsedQty, err = strconv.Atoi(qty)
				check(err)
				gameRound.red = parsedQty
				if parsedQty > minSet.red {
					minSet.red = parsedQty
				}
			case "green":
				var parsedQty, err = strconv.Atoi(qty)
				check(err)
				gameRound.green = parsedQty
				if parsedQty > minSet.green {
					minSet.green = parsedQty
				}
			case "blue":
				var parsedQty, err = strconv.Atoi(qty)
				check(err)
				gameRound.blue = parsedQty
				if parsedQty > minSet.blue {
					minSet.blue = parsedQty
				}
			default:
				fmt.Println("Unknown color: ", color)
			}

		}
		rounds = append(rounds, gameRound)
	}
	var game = newGame(int(gameId), rounds)
	return GameWithMinSet{&game, minSet}
}

func main() {
	f, err := os.Open("../data.txt")
	check(err)
	result := []GameWithMinSet{}
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, parseGame(line))
	}
	sum := 0
	for _, game := range result {
		sum += game.minSet.red * game.minSet.green * game.minSet.blue
	}
	fmt.Println(sum)

	f.Close()
}
