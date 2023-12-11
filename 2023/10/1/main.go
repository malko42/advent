package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Point struct {
	label string
	x     int
	y     int
}

type PipeMap struct {
	pipes [][]string
	start Point
	loop  []Point
}

func (p PipeMap) String() string {
	print := ""
	for _, row := range p.pipes {
		print += fmt.Sprintf("%v\n", row)
	}
	print += fmt.Sprintf("Start: X = %v Y = %v\n", p.start.x, p.start.y)
	print += fmt.Sprintf("Furthest point from start: %v\n", len(p.loop)/2)
	return print
}

func canGoRight(current Point, pipeMap PipeMap) (bool, Point) {
	if current.x < len(pipeMap.pipes[current.y])-1 &&
		(pipeMap.pipes[current.y][current.x+1] == "-" ||
			pipeMap.pipes[current.y][current.x+1] == "J" ||
			pipeMap.pipes[current.y][current.x+1] == "7") &&
		(current.label == "S" ||
			current.label == "-" ||
			current.label == "L" ||
			current.label == "F") {
		return true, Point{pipeMap.pipes[current.y][current.x+1], current.x + 1, current.y}
	}
	return false, Point{"X", -1, -1}
}

func canGoLeft(current Point, pipeMap PipeMap) (bool, Point) {
	if current.x > 0 &&
		(pipeMap.pipes[current.y][current.x-1] == "-" ||
			pipeMap.pipes[current.y][current.x-1] == "F" ||
			pipeMap.pipes[current.y][current.x-1] == "L") &&
		(current.label == "S" ||
			current.label == "-" ||
			current.label == "J" ||
			current.label == "7") {
		return true, Point{pipeMap.pipes[current.y][current.x-1], current.x - 1, current.y}
	}
	return false, Point{"X", -1, -1}
}

func canGoDown(current Point, pipeMap PipeMap) (bool, Point) {
	if current.y < len(pipeMap.pipes)-1 &&
		(pipeMap.pipes[current.y+1][current.x] == "|" ||
			pipeMap.pipes[current.y+1][current.x] == "J" ||
			pipeMap.pipes[current.y+1][current.x] == "L") &&
		(current.label == "S" ||
			current.label == "|" ||
			current.label == "F" ||
			current.label == "7") {
		return true, Point{pipeMap.pipes[current.y+1][current.x], current.x, current.y + 1}
	}
	return false, Point{"X", -1, -1}
}

func canGoUp(current Point, pipeMap PipeMap) (bool, Point) {
	if current.y > 0 &&
		(pipeMap.pipes[current.y-1][current.x] == "|" ||
			pipeMap.pipes[current.y-1][current.x] == "F" ||
			pipeMap.pipes[current.y-1][current.x] == "7") &&
		(current.label == "S" ||
			current.label == "|" ||
			current.label == "J" ||
			current.label == "L") {
		return true, Point{pipeMap.pipes[current.y-1][current.x], current.x, current.y - 1}
	}
	return false, Point{"X", -1, -1}
}

func findNext(current Point, p PipeMap) Point {
	canGo, next := canGoRight(current, p)
	if canGo {
		return next
	}
	canGo, next = canGoDown(current, p)
	if canGo {
		return next
	}
	canGo, next = canGoLeft(current, p)
	if canGo {
		return next
	}
	canGo, next = canGoUp(current, p)
	if canGo {
		return next
	}
	return Point{"X", -1, -1}
}

func followLoop(start Point, p PipeMap) []Point {
	var current = start
	var path []Point
	path = append(path, current)
	for {
		current, p.pipes[current.y][current.x] = findNext(current, p), "X"
		if current.label == "X" {
			break
		}
		path = append(path, current)
	}
	return path

}

func parseData(path string) PipeMap {
	file, err := os.Open(path)
	check(err)
	newMap := PipeMap{}
	lineCount := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var split = strings.Split(line, "")
		if slices.Contains(split, "S") {
			newMap.start = Point{"S", slices.Index(split, "S"), lineCount}
		}
		newMap.pipes = append(newMap.pipes, split)
		lineCount++
	}
	file.Close()
	newMap.loop = followLoop(newMap.start, newMap)
	return newMap
}

func main() {
	result := parseData("../data.txt")
	// result := parseData("../sample.txt")
	fmt.Println(result)
}
