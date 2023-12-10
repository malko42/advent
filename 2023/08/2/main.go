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

func findFirstExit(start string, instructions string, mapData map[string][]string) int {
	counter := 0
	for !isExit(start) {
		for _, instruction := range instructions {
			if instruction == 'L' {
				start = mapData[start][0]
			} else {
				start = mapData[start][1]
			}
			counter++
		}
	}
	return counter
}

func findExit(instructions string, mapData map[string][]string) int {
	var currents = getStartingNodes(mapData)
	stepsToZ := make([]int, len(currents))

	for index, current := range currents {
		stepsToZ[index] = findFirstExit(current, instructions, mapData)
	}

	return LCM(stepsToZ)
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
func LCM(numbers []int) int {
	result := numbers[0] * numbers[1] / GCD(numbers[0], numbers[1])
	for i := 2; i < len(numbers); i++ {
		result = result * numbers[i] / GCD(result, numbers[i])
	}
	return result
}

func main() {
	path, result := parseData("../data.txt")
	// path, result := parseData("sample.txt")
	fmt.Println(findExit(path, result))
}
