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
	seedMap map[int]int
}

func (g Map) getNextLocation(start int) int {
	return g.seedMap[start]
}

func processMapLine(ranges []string, inProgressMap map[int]int) map[int]int {
	origin, err := strconv.Atoi(ranges[0])
	check(err)

	destination, err := strconv.Atoi(ranges[1])
	check(err)

	inProgressMap[origin] = destination

	// bad idea
	// rangeSize, err := strconv.Atoi(ranges[2])
	// check(err)
	//
	// for i := 0; i < rangeSize; i++ {
	// 	inProgressMap[origin+i] = destination+i
	// }

	return inProgressMap
}

func parseMaps(file *os.File) []Map {
	fileScanner := bufio.NewScanner(file)
	result := []Map{}
	var newMap = Map{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var split = strings.Fields(line)
		fmt.Println(len(split))
		// start of a map
		if len(split) == 2 {
			newMap.label = split[0]
			newMap.seedMap = make(map[int]int)

		}
		if len(split) == 3 {
			newMap.seedMap = processMapLine(split, newMap.seedMap)
		}
		// End of a map
		if len(split) == 0 {
			result = append(result, newMap)
			newMap = Map{}
		}

		fmt.Println(split)
	}
	file.Close()
	return result
}

func browseMaps(maps []Map, seed string) int {
intseed, err := strconv.Atoi(seed)
check(err)
location := 0
// check si seed est entre start et start + range et si oui on return seed - start + destination
	for _, currentMap := range maps {



	}

}

func main() {
	seeds := strings.Fields("1187290020 247767461 40283135 64738286 2044483296 66221787 1777809491 103070898 108732160 261552692 3810626561 257826205 3045614911 65672948 744199732 300163578 3438684365 82800966 2808575117 229295075")
	fmt.Println(seeds)
	f, err := os.Open("../data.txt")
	check(err)
	result := []Map{}
	result = parseMaps(f)
	fmt.Println(result)
}
