package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func isNumber(r string) (bool, int) {
	num, err := strconv.Atoi(r)
	if err == nil {
		return true, num
	}
	return false, 0
}

func isSymbol(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return false
	}
	if s == "." {
		return false
	}
	return true
}

func hasAdjacentSymbol(x int, y int, data [][]string) bool {
	if x > 0 && y > 0 && x < len(data[y])-1 && y < len(data)-1 {
		if isSymbol(data[y][x-1]) ||
			isSymbol(data[y-1][x-1]) ||
			isSymbol(data[y-1][x]) ||
			isSymbol(data[y-1][x+1]) ||
			isSymbol(data[y][x+1]) ||
			isSymbol(data[y+1][x+1]) ||
			isSymbol(data[y+1][x]) ||
			isSymbol(data[y+1][x-1]) {
			return true
		}
		return false
	}

	if x == 0 && y == 0 {
		if isSymbol(data[y][x+1]) ||
			isSymbol(data[y+1][x+1]) ||
			isSymbol(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == 0 && y == len(data)-1 {
		if isSymbol(data[y][x+1]) ||
			isSymbol(data[y-1][x+1]) ||
			isSymbol(data[y-1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y == 0 {
		if isSymbol(data[y][x-1]) ||
			isSymbol(data[y+1][x-1]) ||
			isSymbol(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y == len(data)-1 {
		if isSymbol(data[y][x-1]) ||
			isSymbol(data[y-1][x-1]) ||
			isSymbol(data[y-1][x]) {
			return true
		}
		return false
	}

	if y == 0 && x > 0 && x < len(data[y])-1 {
		if isSymbol(data[y][x-1]) ||
			isSymbol(data[y+1][x-1]) ||
			isSymbol(data[y+1][x]) ||
			isSymbol(data[y+1][x+1]) ||
			isSymbol(data[y][x+1]) {
			return true
		}
		return false
	}

	if y == len(data)-1 && x > 0 && x < len(data[y])-1 {
		if isSymbol(data[y][x-1]) ||
			isSymbol(data[y-1][x-1]) ||
			isSymbol(data[y-1][x]) ||
			isSymbol(data[y-1][x+1]) ||
			isSymbol(data[y][x+1]) {
			return true
		}
		return false
	}

	if x == 0 && y > 0 && y < len(data)-1 {
		if isSymbol(data[y-1][x]) ||
			isSymbol(data[y-1][x+1]) ||
			isSymbol(data[y][x+1]) ||
			isSymbol(data[y+1][x+1]) ||
			isSymbol(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y > 0 && y < len(data)-1 {
		if isSymbol(data[y-1][x]) ||
			isSymbol(data[y-1][x-1]) ||
			isSymbol(data[y][x-1]) ||
			isSymbol(data[y+1][x-1]) ||
			isSymbol(data[y+1][x]) {
			return true
		}
		return false
	}
	return false
}

func parseData(f *bufio.Scanner) ([][]string, error) {
	result := [][]string{}

	for f.Scan() {
		row := []string{}
		line := f.Text()
		for _, char := range line {
			row = append(row, string(char))
		}
		result = append(result, row)
	}
	return result, nil
}

func getIntFromArray(table []int) int {
	result := 0
	for i, num := range table {
		result += num * int(math.Pow(10, float64(len(table)-i-1)))
	}
	return result
}

func main() {
	f, err := os.Open("../data.txt")
	check(err)
	fileScanner := bufio.NewScanner(f)
	table, err := parseData(fileScanner)
	check(err)
	validNumbers := []int{}
	for y, row := range table {
		pile := []int{}
		symbolHasBeenEncountered := false
		for x, char := range row {
			var isDigit, value = isNumber(char)
			if isDigit {
				pile = append(pile, value)
			}
			if isDigit && hasAdjacentSymbol(x, y, table) {
				symbolHasBeenEncountered = true
			}
			// end of a number and no adjacent symbol
			if char == "." && !symbolHasBeenEncountered {
				pile = []int{}
				symbolHasBeenEncountered = false
			}
			// end of a number whether it's a symbol or end of line
			if (isDigit && x == len(row)-1 && symbolHasBeenEncountered) ||
				(!isDigit && symbolHasBeenEncountered) {
				intfromarray := getIntFromArray(pile)
				if intfromarray != 0 {
					validNumbers = append(validNumbers, intfromarray)

					pile = []int{}
					symbolHasBeenEncountered = false
				}
			}
		}
	}

	sum := 0
	for _, num := range validNumbers {
		sum += num
	}

	fmt.Println(sum)

	f.Close()
}
