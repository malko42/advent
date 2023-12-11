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

type Tile struct {
	label          string
	x              int
	y              int
	hasBeenVisited bool
}

type PipeMap struct {
	tiles [][]Tile
	start *Tile
	loop  []Tile
}

func (p PipeMap) String() string {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
	print := ""
	for _, row := range p.tiles {
		print += "["
		for _, tile := range row {
			if tile.hasBeenVisited {
				print += colorRed
			} else {
				print += colorNone
			}
			print += fmt.Sprintf("%v ", tile.label)
		}
		print += "]\n"
	}
	print += fmt.Sprintf("Start: X = %v Y = %v\n", p.start.x, p.start.y)
	print += fmt.Sprintf("Furthest point from start: %v\n", len(p.loop)/2)
	return print
}

func (t Tile) isStart() bool {
	return t.label == "S"
}

func canGoRight(current Tile, pipeMap PipeMap) bool {
	next := pipeMap.tiles[current.y][current.x+1]
	return current.x < len(pipeMap.tiles[current.y])-1 &&
	!next.hasBeenVisited &&
			(next.label == "-" ||
				next.label == "J" ||
				next.label == "7") &&
			(current.label == "S" ||
				current.label == "-" ||
				current.label == "L" ||
				current.label == "F")
}

func canGoLeft(current Tile, pipeMap PipeMap) bool {
	next := pipeMap.tiles[current.y][current.x-1]
	return current.x > 0 &&
	!next.hasBeenVisited &&
			(next.label == "-" ||
				next.label == "F" ||
				next.label == "L") &&
			(current.label == "S" ||
				current.label == "-" ||
				current.label == "J" ||
				current.label == "7")
}

func canGoDown(current Tile, pipeMap PipeMap) bool {
	next := pipeMap.tiles[current.y+1][current.x]
	return current.y < len(pipeMap.tiles)-1 &&
	!next.hasBeenVisited &&
			(next.label == "|" ||
				next.label == "J" ||
				next.label == "L") &&
			(current.label == "S" ||
				current.label == "|" ||
				current.label == "F" ||
				current.label == "7")
}

func canGoUp(current Tile, pipeMap PipeMap) bool {
	next := pipeMap.tiles[current.y-1][current.x]
	return current.y > 0 &&
	!next.hasBeenVisited &&
			(next.label == "|" ||
				next.label == "F" ||
				next.label == "7") &&
			(current.label == "S" ||
				current.label == "|" ||
				current.label == "J" ||
				current.label == "L")
}

func findNext(current Tile, p *PipeMap) Tile {

	if canGoRight(current, *p) {
		p.tiles[current.y][current.x].hasBeenVisited = true
		return p.tiles[current.y][current.x+1]
	}
	if canGoDown(current, *p) {
		fmt.Println("Going down!")
		p.tiles[current.y][current.x].hasBeenVisited = true
		return p.tiles[current.y+1][current.x]
	}
	if canGoLeft(current, *p) {
		fmt.Println("Going left!")
		p.tiles[current.y][current.x].hasBeenVisited = true
		return p.tiles[current.y][current.x-1]
	}
	if canGoUp(current, *p) {
		fmt.Println("Going up!")
		p.tiles[current.y][current.x].hasBeenVisited = true
		return p.tiles[current.y-1][current.x]
	}
	fmt.Println("Going nowhere")
	return Tile{"X", -1, -1, false}
}

func followLoop(p *PipeMap) []Tile {
	var current = *p.start
	var path []Tile
	path = append(path, current)
	for {
		current = findNext(current, p)
		fmt.Println(current)
		if current.isStart(){
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
		lineOfTiles := []Tile{}
		for i, char := range split {
			split[i] = char
			newTile := Tile{char, i, lineCount, false}
			if char == "S" {
				newMap.start = &newTile
			}
			lineOfTiles = append(lineOfTiles, newTile)
		}
		newMap.tiles = append(newMap.tiles, lineOfTiles)
		lineCount++
	}
	file.Close()
	newMap.loop = followLoop(&newMap)
	return newMap
}

func main() {
	result := parseData("../data.txt")
	// result := parseData("../sample.txt")
	fmt.Println(result)
}

