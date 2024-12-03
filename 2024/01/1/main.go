package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var log = fmt.Println

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type IdsList struct {
	left  []int
	right []int
}

type IdsCouple struct {
	left  int
	right int
	diff  int
}

func parseData(path string) IdsList {
	file, err := os.Open(path)
	check(err)
	idsList := IdsList{}
	lineCount := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var split = strings.Split(line, "   ")
		leftId, err := strconv.Atoi(split[0])
		check(err)
		rightId, err := strconv.Atoi(split[1])
		check(err)
		idsList.left = append(idsList.left, leftId)
		idsList.right = append(idsList.right, rightId)
		lineCount++
	}
	file.Close()
	return idsList
}

func sortIdsLists(idsList IdsList) IdsList {
	sort.Slice(idsList.left, func(i, j int) bool {
		return idsList.left[i] < idsList.left[j]
	})
	sort.Slice(idsList.right, func(i, j int) bool {
		return idsList.right[i] < idsList.right[j]
	})
	return idsList
}

func main() {
	parsedData := parseData("../data.txt")
	parsedData = sortIdsLists(parsedData)
	// result := parseData("../sample.txt")
	fmt.Println(parsedData)
	// log(len(result.left))
	// log(len(result.right))
	result := []IdsCouple{}
	for i := 0; i < len(parsedData.left); i++ {
		result = append(result, IdsCouple{parsedData.left[i], parsedData.right[i], int(math.Abs(
			float64(parsedData.left[i] - parsedData.right[i])))})
	}
	add := 0
	for _, couple := range result {
		add += couple.diff
		fmt.Println(couple)
	}
	log(add)
}
