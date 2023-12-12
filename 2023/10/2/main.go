package main

import (
	"bufio"
	"errors"
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
	const colorYellow = "\033[1;33m"
	const colorNone = "\033[0m"
	max := p.loop[len(p.loop)/2]

	print := ""
	for _, row := range p.tiles {
		print += "["
		for _, tile := range row {
			if tile.hasBeenVisited {
				if tile.x == max.x && tile.y == max.y {
					print += colorYellow
				} else {
					print += colorRed
				}
			} else {
				print += colorNone
			}
			print += fmt.Sprintf("%v ", tile.label)
		}
		print += "]\n"
	}
	print += fmt.Sprintf("Start: X = %v Y = %v\n", p.start.x, p.start.y)
	print += fmt.Sprintf("Furthest point from start: X = %v Y= %v at index %v\n", max.x, max.y, len(p.loop)/2)
	return print
}

func (t Tile) isStart() bool {
	return t.label == "S"
}

func canGoRight(current Tile, pipeMap PipeMap) bool {
	return current.x < len(pipeMap.tiles[current.y])-1 &&
		(!pipeMap.tiles[current.y][current.x+1].hasBeenVisited ||
			(pipeMap.tiles[current.y][current.x+1].hasBeenVisited && pipeMap.tiles[current.y][current.x+1].isStart())) &&
		(pipeMap.tiles[current.y][current.x+1].label == "-" ||
			pipeMap.tiles[current.y][current.x+1].label == "J" ||
			pipeMap.tiles[current.y][current.x+1].label == "S" ||
			pipeMap.tiles[current.y][current.x+1].label == "7") &&
		(current.label == "S" ||
			current.label == "-" ||
			current.label == "L" ||
			current.label == "F")
}

func canGoLeft(current Tile, pipeMap PipeMap) bool {
	return current.x > 0 &&
		(!pipeMap.tiles[current.y][current.x-1].hasBeenVisited || (pipeMap.tiles[current.y][current.x-1].hasBeenVisited && pipeMap.tiles[current.y][current.x-1].isStart())) &&
		(pipeMap.tiles[current.y][current.x-1].label == "-" ||
			pipeMap.tiles[current.y][current.x-1].label == "F" ||
			pipeMap.tiles[current.y][current.x-1].label == "S" ||
			pipeMap.tiles[current.y][current.x-1].label == "L") &&
		(current.label == "S" ||
			current.label == "-" ||
			current.label == "J" ||
			current.label == "7")
}

func canGoDown(current Tile, pipeMap PipeMap) bool {
	return current.y < len(pipeMap.tiles)-1 &&
		(!pipeMap.tiles[current.y+1][current.x].hasBeenVisited || (pipeMap.tiles[current.y+1][current.x].hasBeenVisited && pipeMap.tiles[current.y+1][current.x].isStart())) &&
		(pipeMap.tiles[current.y+1][current.x].label == "|" ||
			pipeMap.tiles[current.y+1][current.x].label == "J" ||
			pipeMap.tiles[current.y+1][current.x].label == "S" ||
			pipeMap.tiles[current.y+1][current.x].label == "L") &&
		(current.label == "S" ||
			current.label == "|" ||
			current.label == "F" ||
			current.label == "7")
}

func canGoUp(current Tile, pipeMap PipeMap) bool {
	return current.y > 0 &&
		(!pipeMap.tiles[current.y-1][current.x].hasBeenVisited || (pipeMap.tiles[current.y-1][current.x].hasBeenVisited && pipeMap.tiles[current.y-1][current.x].isStart())) &&
		(pipeMap.tiles[current.y-1][current.x].label == "|" ||
			pipeMap.tiles[current.y-1][current.x].label == "F" ||
			pipeMap.tiles[current.y-1][current.x].label == "S" ||
			pipeMap.tiles[current.y-1][current.x].label == "7") &&
		(current.label == "S" ||
			current.label == "|" ||
			current.label == "J" ||
			current.label == "L")
}

func findNext(current Tile, p *PipeMap) (Tile, error) {
	if canGoRight(current, *p) {
		p.tiles[current.y][current.x].hasBeenVisited = true
		p.tiles[current.y][current.x] = getPrettyTile(p.tiles[current.y][current.x])
		return p.tiles[current.y][current.x+1], nil
	}
	if canGoDown(current, *p) {
		p.tiles[current.y][current.x].hasBeenVisited = true
		p.tiles[current.y][current.x] = getPrettyTile(p.tiles[current.y][current.x])
		return p.tiles[current.y+1][current.x], nil
	}
	if canGoLeft(current, *p) {
		p.tiles[current.y][current.x].hasBeenVisited = true
		p.tiles[current.y][current.x] = getPrettyTile(p.tiles[current.y][current.x])
		return p.tiles[current.y][current.x-1], nil
	}
	if canGoUp(current, *p) {
		p.tiles[current.y][current.x].hasBeenVisited = true
		p.tiles[current.y][current.x] = getPrettyTile(p.tiles[current.y][current.x])
		return p.tiles[current.y-1][current.x], nil
	}
	return Tile{"X", -1, -1, false}, errors.New("No valid path found")
}

func followLoop(p *PipeMap) []Tile {
	var current = *p.start
	var path []Tile
	path = append(path, current)
	for {
		var err error
		current, err = findNext(current, p)
		check(err)

		if current.isStart() {
			break
		}
		path = append(path, current)
	}
	return path

}

func getPrettyTile(tile Tile) Tile {
	if tile.label == "L" {
		tile.label = "└"
	}
	if tile.label == "F" {
		tile.label = "┌"
	}
	if tile.label == "J" {
		tile.label = "┘"
	}
	if tile.label == "7" {
		tile.label = "┐"
	}
	if tile.label == "|" {
		tile.label = "│"
	}
	if tile.label == "-" {
		tile.label = "─"
	}
	return tile
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
