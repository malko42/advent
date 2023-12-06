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

type Data struct {
	label   string
	numbers []int
}

func parseData(path string) []Data {
	file, err := os.Open(path)
	check(err)
	result := []Data{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		newData := Data{}
		numbers := []int{}
		var split = strings.Fields(line)
		for _, stringNumber := range split {
			number, err := strconv.Atoi(stringNumber)
			check(err)
			numbers = append(numbers, number)
		}

		newData.label = "lol"
		newData.numbers = numbers

		result = append(result, newData)
	}
	file.Close()
	return result
}

func main() {
	result := []Data{}
	result = parseData("../data.txt")
	for _, currentData := range result {
		fmt.Println(currentData)
	}
}
