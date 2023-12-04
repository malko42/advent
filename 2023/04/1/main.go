package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Game struct {
	id             int
	winningNumbers []int
	givenNumbers   []int
	commonNumbers  []int
}

func (g Game) isWon() bool {
	return len(g.commonNumbers) > 0
}

func (g Game) getScore() int {
	if len(g.commonNumbers) > 0 {
		return int(math.Pow(2, float64(len(g.commonNumbers)-1)))
	}
	return 0
}

func parseGame(line string) Game {
	var split = strings.Split(line, ":")
	var gameString = split[0]
	var numbersString = split[1]
	gameId, err := strconv.Atoi(strings.Fields(gameString)[1])
	check(err)
	winningNumbers := []int{}
	numbers := []int{}
	commonNumbers := []int{}
	var splitNumbersString = strings.Split(numbersString, "|")
	var winningNumbersString = splitNumbersString[0]
	var givenNumbersString = splitNumbersString[1]

	for _, winningNumber := range strings.Fields(winningNumbersString) {
		winningNumberInt, err := strconv.Atoi(winningNumber)
		check(err)
		winningNumbers = append(winningNumbers, winningNumberInt)
	}

	for _, givenNumber := range strings.Fields(givenNumbersString) {
		givenNumberInt, err := strconv.Atoi(givenNumber)
		check(err)
		numbers = append(numbers, givenNumberInt)
	}

	for _, winningNumber := range winningNumbers {
		for _, givenNumber := range numbers {
			if winningNumber == givenNumber {
				commonNumbers = append(commonNumbers, winningNumber)
			}
		}
	}

	return Game{gameId, winningNumbers, numbers, commonNumbers}
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
	sum := 0
	for _, game := range result {
		if game.isWon() {
			sum += game.getScore()
		}
	}
	fmt.Println(sum)

	f.Close()
}
