package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func isExit(label string) bool {
	return string(label[2]) == "Z"
}

func isStart(label string) bool {
	return string(label[2]) == "A"
}

func hasReachedExit(currents []string) bool {
	// every current is an exit
	status := true
	for _, current := range currents {
		status = status && isExit(current)
	}
	return status
}

func findExit(instructions string, mapData map[string][]string) int {
	var currents = getStartingNodes(mapData)
	counter := 0
	fastestToExit := make([]int, len(currents))
	for !hasReachedExit(currents) {
		path := strings.Split(instructions, "")
		for _, instruction := range path {
			counter++
			for index, current := range currents {
				if isExit(current) {
					if fastestToExit[index] == 0 {
						fastestToExit[index] = counter
					}
				}
				if instruction == "L" {
					currents[index] = mapData[current][0]
				}
				if instruction == "R" {
					currents[index] = mapData[current][1]
				}
			}
		}
	}
	return counter
}

func getStartingNodes(mapData map[string][]string) []string {
	var result []string
	for key := range mapData {
		if isStart(key) {
			result = append(result, key)
		}
	}
	return result
}

func parseData(path string) (lrPath string, result map[string][]string) {
	file, err := os.Open(path)
	check(err)
	result = map[string][]string{}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()
	lrPath = fileScanner.Text()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		var split = strings.Fields(line)
		// example of split
		// XGH
		// =
		// (JQN,
		// TVG)
		result[string(split[0])] = []string{string(split[2])[1 : len(split[2])-1], string(split[3][0 : len(split[3])-1])}
	}
	file.Close()
	return lrPath, result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	path, result := parseData("../data.txt")
	// path, result := parseData("sample.txt")
	fmt.Println(findExit(path, result))
	// fmt.Println(LCM(11310, 20777, 17622, 13940, 15518, 20778))
}
