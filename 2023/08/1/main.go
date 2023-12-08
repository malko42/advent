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

type Node struct {
	label string
	left  *Node
	right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("label: %s, left: %s, right: %s", n.label, n.left, n.right)
}

func findExit(instructions string, mapData map[string][]string, start string) int {
	var current = start
	counter := 0
	for current != "ZZZ" {
		path := strings.Split(instructions, "")
		for _, instruction := range path {
			counter++
			if instruction == "L" {
				current = mapData[current][0]
			}
			if instruction == "R" {
				current = mapData[current][1]
			}
		}
	}
	return counter
}

var log = fmt.Println

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

func main() {
	path, result := parseData("../data.txt")
	// path, result := parseData("../sample.txt")
	fmt.Println(findExit(path, result, "AAA"))
}
