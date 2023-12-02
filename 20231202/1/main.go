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

func newRound(red int, green int, blue int) Round {
	return Round{red, green, blue}
}

func newGame(id int, rounds []Round) Game {
	return Game{id, rounds}
}

func filterGames(games []Game, maxRed int, maxGreen int, maxBlue int) []Game {
	var result = []Game{}
gameloop:
	for _, game := range games {
		for _, round := range game.rounds {
			if round.red > maxRed || round.green > maxGreen || round.blue > maxBlue {
				continue gameloop
			}
		}
		result = append(result, game)
	}
	return result
}

func parseGame(line string) Game {
	var split = strings.Split(line, ":")
	var gameString = split[0]
	var roundsString = split[1]
	gameId, err := strconv.Atoi(strings.Split(gameString, " ")[1])
	check(err)
	var rounds = []Round{}
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
				gameRound.red, err = strconv.Atoi(qty)
				check(err)
			case "green":
				gameRound.green, err = strconv.Atoi(qty)
				check(err)
			case "blue":
				gameRound.blue, err = strconv.Atoi(qty)
				check(err)
			default:
				fmt.Println("Unknown color: ", color)
			}

		}
		rounds = append(rounds, gameRound)
	}
	return newGame(int(gameId), rounds)
}

func main() {
	f, err := os.Open("../data.txt")
	check(err)
	result := []Game{}
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, parseGame(line))
	}
	var filteredGames = filterGames(result, 12, 13, 14)
	sum := 0
	for _, game := range filteredGames {
		sum += game.id
	}

	fmt.Println(sum)

	f.Close()
}
