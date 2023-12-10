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
	recording  []int
	regression [][]int
}

func (data Data) extrapolateNextReading() int {
	result := 0
	for i := len(data.regression) - 1; i >= 0; i-- {
		result += data.regression[i][len(data.regression[i])-1]
	}
	return result + data.recording[len(data.recording)-1]
}

func isRegressingOver(recording []int) bool {
	status := true
	for _, number := range recording {
		status = status && number == 0
	}
	return status
}

func getNextRegression(recording []int) []int {
	result := []int{}
	for i := 0; i < len(recording)-1; i++ {
		result = append(result, recording[i+1]-recording[i])
	}

	return result
}

func getRegressions(recording []int) [][]int {
	result := [][]int{}
	for !isRegressingOver(recording) {
		recording = getNextRegression(recording)
		result = append(result, recording)
	}

	return result
}

func parseData(path string) []Data {
	file, err := os.Open(path)
	check(err)
	result := []Data{}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		newData := Data{}
		recording := []int{}
		var split = strings.Fields(line)
		for _, stringNumber := range split {
			number, err := strconv.Atoi(stringNumber)
			check(err)
			recording = append(recording, number)
		}

		newData.recording = recording
		newData.regression = getRegressions(recording)

		result = append(result, newData)
	}
	file.Close()
	return result
}

func main() {
	result := []Data{}
	result = parseData("../data.txt")
	// result = parseData("../sample.txt")
	sum := 0
	for _, currentData := range result {
		sum += currentData.extrapolateNextReading()
	}
	fmt.Println(sum)
}
