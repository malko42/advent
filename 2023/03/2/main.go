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

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	}
	return false
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

func hasAdjacentDigit(x int, y int, data [][]string) bool {
	if x > 0 && y > 0 && x < len(data[y])-1 && y < len(data)-1 {
		if isDigit(data[y][x-1]) ||
			isDigit(data[y-1][x-1]) ||
			isDigit(data[y-1][x]) ||
			isDigit(data[y-1][x+1]) ||
			isDigit(data[y][x+1]) ||
			isDigit(data[y+1][x+1]) ||
			isDigit(data[y+1][x]) ||
			isDigit(data[y+1][x-1]) {
			return true
		}
		return false
	}

	if x == 0 && y == 0 {
		if isDigit(data[y][x+1]) ||
			isDigit(data[y+1][x+1]) ||
			isDigit(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == 0 && y == len(data)-1 {
		if isDigit(data[y][x+1]) ||
			isDigit(data[y-1][x+1]) ||
			isDigit(data[y-1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y == 0 {
		if isDigit(data[y][x-1]) ||
			isDigit(data[y+1][x-1]) ||
			isDigit(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y == len(data)-1 {
		if isDigit(data[y][x-1]) ||
			isDigit(data[y-1][x-1]) ||
			isDigit(data[y-1][x]) {
			return true
		}
		return false
	}

	if y == 0 && x > 0 && x < len(data[y])-1 {
		if isDigit(data[y][x-1]) ||
			isDigit(data[y+1][x-1]) ||
			isDigit(data[y+1][x]) ||
			isDigit(data[y+1][x+1]) ||
			isDigit(data[y][x+1]) {
			return true
		}
		return false
	}

	if y == len(data)-1 && x > 0 && x < len(data[y])-1 {
		if isDigit(data[y][x-1]) ||
			isDigit(data[y-1][x-1]) ||
			isDigit(data[y-1][x]) ||
			isDigit(data[y-1][x+1]) ||
			isDigit(data[y][x+1]) {
			return true
		}
		return false
	}

	if x == 0 && y > 0 && y < len(data)-1 {
		if isDigit(data[y-1][x]) ||
			isDigit(data[y-1][x+1]) ||
			isDigit(data[y][x+1]) ||
			isDigit(data[y+1][x+1]) ||
			isDigit(data[y+1][x]) {
			return true
		}
		return false
	}

	if x == len(data[y])-1 && y > 0 && y < len(data)-1 {
		if isDigit(data[y-1][x]) ||
			isDigit(data[y-1][x-1]) ||
			isDigit(data[y][x-1]) ||
			isDigit(data[y+1][x-1]) ||
			isDigit(data[y+1][x]) {
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

func simplifyData(data [][]string) [][]string {
	result := [][]string{}
	for _, row := range data {
		newRow := []string{}
		for _, char := range row {
			if isDigit(string(char)) {
				newRow = append(newRow, string(char))
				continue
			}
			if string(char) != "*" {
				newRow = append(newRow, ".")
				continue
			}
			newRow = append(newRow, char)
		}
		result = append(result, newRow)
	}
	return result
}

func getIntFromArray(table []int) int {
	result := 0
	for i, num := range table {
		result += num * int(math.Pow(10, float64(len(table)-i-1)))
	}
	return result
}

func getAdjacentNumbers(x int, y int, data [][]string) []int {
	result := []int{}

		if isDigit(data[y][x-1]) {
			tmp := ""
			end := x-1
			for data[y][end] != "." && end > 0{
				tmp = data[y][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}
		if isDigit(data[y-1][x-1]) && !isDigit(data[y-1][x]) {
			tmp := ""
			end := x-1
			for data[y-1][end] != "." && end > 0 {
				tmp = data[y-1][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}


		if isDigit(data[y-1][x-1]) && isDigit(data[y-1][x]) && !isDigit(data[y-1][x+1]) {
			tmp := ""
			end := x
			for data[y-1][end] != "." && end > 0{
				tmp = data[y-1][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}
		
		if isDigit(data[y-1][x-1]) && isDigit(data[y-1][x]) && isDigit(data[y-1][x+1]) {
			tmp := ""
			end := x+1
			for data[y-1][end] != "." && end > 0{
				tmp = data[y-1][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y-1][x]) && !isDigit(data[y-1][x-1]) {
			tmp := ""
			start:= x
			for data[y-1][start] != "." && start < len(data[y-1])-1{
				tmp += data[y-1][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if (isDigit(data[y-1][x+1])) && !isDigit(data[y-1][x]) {
			tmp := ""
			start:= x+1
			for data[y-1][start] != "." && start < len(data[y-1])-1{
				tmp += data[y-1][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y][x+1]) {
			tmp := ""
			start:= x+1
			for data[y][start] != "." && start < len(data[y])-1{
				tmp += data[y][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y+1][x+1]) && !isDigit(data[y+1][x]) {
			tmp := ""
			start:= x+1
			for data[y+1][start] != "." && start < len(data[y+1])-1{
				tmp += data[y+1][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y+1][x+1]) && isDigit(data[y+1][x]) && !isDigit(data[y+1][x-1]) {
			tmp := ""
			start:= x
			for data[y+1][start] != "." && start < len(data[y+1])-1{
				tmp += data[y+1][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y+1][x+1]) && isDigit(data[y+1][x]) && isDigit(data[y+1][x-1]) {
			tmp := ""
			start:= x-1
			for data[y+1][start] != "." && start < len(data[y+1])-1{
				tmp += data[y+1][start]
				start++
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}

		if isDigit(data[y+1][x]) && !isDigit(data[y+1][x+1]) {
			tmp := ""
			end := x
			for data[y+1][end] != "." && end > 0 {
				tmp = data[y+1][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}
		if isDigit(data[y+1][x-1]) && !isDigit(data[y+1][x]) {
			tmp := ""
			end := x-1
			for data[y+1][end] != "." && end > 0{
				tmp = data[y+1][end] + tmp
				end--
			}
			num, _ := strconv.Atoi(tmp)
			result = append(result, num)
		}


		return result
}

func main() {
	f, err := os.Open("../data.txt")
	check(err)
	fileScanner := bufio.NewScanner(f)
	table, err := parseData(fileScanner)
	check(err)
	table = simplifyData(table)
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
				for i := 0; i < len(pile); i++ {
					table[y][x-i-1] = "."
				}

				pile = []int{}
				symbolHasBeenEncountered = false
			}
			if (x == len(row)-1 && !symbolHasBeenEncountered) {

				for i := 0; i < len(pile); i++ {
					table[y][x-i] = "."
				}

				pile = []int{}
				symbolHasBeenEncountered = false
			}
			// end of a number whether it's a symbol or end of line
			if (isDigit && x == len(row)-1 && symbolHasBeenEncountered) ||
				(!isDigit && symbolHasBeenEncountered) {
				intfromarray := getIntFromArray(pile)
				if intfromarray != 0 {

					pile = []int{}
					symbolHasBeenEncountered = false
				}
			}
		}
	}

	result := [][]int{}

	for y, row := range table {
		for x, char := range row {
			if (char == "*") {
				numbers := getAdjacentNumbers(x, y, table)
				result = append(result, numbers)
			}
		}
	}
	for _, row := range table {
		fmt.Println(row)
	}

	// fmt.Println(result)


	for _, gears := range result {
		if len(gears) == 2 {
			validNumbers = append(validNumbers, gears[0] * gears[1])
		}
	}


	sum := 0
	for _, num := range validNumbers {
		sum += num
	}
	fmt.Println(sum)


	f.Close()
}
