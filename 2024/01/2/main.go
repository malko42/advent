package main

import (
	"bufio"
	"fmt"
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
	// log(len(parsedData.left))
	// log(len(parsedData.right))
	occurenceMap := make(map[int]int)
	for i := 0; i < len(parsedData.left); i++ {
		for j := 0; j < len(parsedData.right); j++ {
			if parsedData.left[i] < parsedData.right[j] {
				break
			}
			if parsedData.left[i] == parsedData.right[j] {
				occurenceMap[parsedData.left[i]]++
			}
		}
	}

	log(occurenceMap)
	result := 0
	for value, count := range occurenceMap {
		result += value * count
	}
	log(result)
}

