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

type Map struct {
	label   string
	seedMap []Route
}

type Route struct {
	source      int
	destination int
	rangeSize   int
}

type Seed struct {
	start     int
	rangeSize int
}

func processRoute(ranges []string) Route {
	origin, err := strconv.Atoi(ranges[1])
	check(err)

	destination, err := strconv.Atoi(ranges[0])
	check(err)

	rangeSize, err := strconv.Atoi(ranges[2])
	check(err)

	return Route{origin, destination, rangeSize}
}

func parseMaps(file *os.File) []Map {
	fileScanner := bufio.NewScanner(file)
	maps := []Map{}
	newMap := Map{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var split = strings.Fields(line)
		// start of a map
		if len(split) == 2 {
			newMap.label = split[0]
			newMap.seedMap = []Route{}
		}
		if len(split) == 3 {
			newMap.seedMap = append(newMap.seedMap, processRoute(split))
		}
		// End of a map
		if len(split) == 0 {
			maps = append(maps, newMap)
			newMap = Map{}
		}

	}
	file.Close()
	for _, currentMap := range maps {
		for i := 0; i < len(currentMap.seedMap); i++ {
			for j := i + 1; j < len(currentMap.seedMap); j++ {
				if currentMap.seedMap[i].source > currentMap.seedMap[j].source {
					currentMap.seedMap[i], currentMap.seedMap[j] = currentMap.seedMap[j], currentMap.seedMap[i]
				}
			}
		}
	}
	return maps
}

func browseMaps(maps []Map, seed Seed) int {
	seedInt := seed.start
	for _, currentMap := range maps {
		for _, currentRoute := range currentMap.seedMap {
			if seedInt >= currentRoute.source && seedInt < currentRoute.source+currentRoute.rangeSize {
				seedInt = (seedInt - currentRoute.source) + currentRoute.destination
				break
			}
		}
	}

	return seedInt
}

func parseSeeds(seedsString string) []Seed {
	ints := strings.Fields(seedsString)
	result := []Seed{}

	for i := 0; i < len(ints); i = i + 2 {
		parsedSeed, err := strconv.Atoi(ints[i])
		check(err)
		parsedRange, err := strconv.Atoi(ints[i+1])
		newSeed := Seed{parsedSeed, parsedRange}
		result = append(result, newSeed)
	}
	return result
}

func main() {
	seeds := parseSeeds("1187290020 247767461 40283135 64738286 2044483296 66221787 1777809491 103070898 108732160 261552692 3810626561 257826205 3045614911 65672948 744199732 300163578 3438684365 82800966 2808575117 229295075")
	f, err := os.Open("../data.txt")
	check(err)
	result := []Map{}
	result = parseMaps(f)
	lowestLocation := 0

	for _, seed := range seeds {
		lowestLocation = browseMaps(result, seed)
		for i := seed.start; i < seed.start+seed.rangeSize-1; i++ {
			newSeed := Seed{i, 0}
			endLocation := browseMaps(result, newSeed)
			if endLocation < lowestLocation {
				lowestLocation = endLocation
			}
		}
	}
	fmt.Println(lowestLocation)
}
