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

type Scenario struct {
	buttonPushed int
	distanceRan  int
}

type Race struct {
	timeAllowed     int
	distanceToRun   int
	winningScenarii []Scenario
}

func parseData(path string) []Race {
	file, err := os.Open(path)
	check(err)
	races := []Race{}
	fileScanner := bufio.NewScanner(file)
	newRace := Race{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")
		if split[0] == "Time" {
			concatenated := strings.Join(strings.Fields(split[1]), "")
			newRace.timeAllowed, err = strconv.Atoi(concatenated)
			check(err)
		}
		if split[0] == "Distance" {
			concatenated := strings.Join(strings.Fields(split[1]), "")
			newRace.distanceToRun, err = strconv.Atoi(concatenated)
			check(err)
		}
	}
	races = append(races, newRace)
	file.Close()
	return races
}

func (r *Race) getWinningScenarii() []Scenario {
	const SPEED = 1
	scenarii := []Scenario{}
	for i := 0; i < r.timeAllowed; i++ {
		currentSpeed := SPEED * i
		timeRemaining := r.timeAllowed - i
		distanceRange := currentSpeed * timeRemaining
		if distanceRange >= r.distanceToRun {
			scenarii = append(scenarii, Scenario{buttonPushed: i, distanceRan: distanceRange})
		}
	}
	return scenarii
}

func main() {
	result := []Race{}
	result = parseData("../data.txt")
	factor := 1
	for _, currentData := range result {
		currentData.winningScenarii = currentData.getWinningScenarii()
		factor *= len(currentData.winningScenarii)
	}
	fmt.Println(factor)
}
