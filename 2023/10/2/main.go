package main

import (
	"bufio"
	"errors"
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

type Tile struct {
	label          string
	x              int
	y              int
	hasBeenVisited bool
	group          int
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
	const colorGreen = "\033[0;32m"
	const colorBlue = "\033[0;34m"
	const colorCyan = "\033[0;36m"
	const colorPurple = "\033[0;35m"
	const colorOrange = "\033[0;33m"
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
				if tile.group == 1 {
					print += colorGreen
				} else if tile.group == 2 {
					print += colorPurple
				} else {
					print += colorNone
				}
			}
			print += fmt.Sprintf("%v ", tile.label)
			print += colorNone
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
	return Tile{"X", -1, -1, false, 0}, errors.New("No valid path found")
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

func getLeftHandTile(pipeMap PipeMap, prev Tile, current Tile) (tiles []Tile, err error) {
	xDirection := current.x - prev.x
	yDirection := current.y - prev.y
	if xDirection == 1 && current.y > 0 {
		return []Tile{pipeMap.tiles[current.y-1][current.x]}, nil
	}
	if xDirection == -1 && current.y < len(pipeMap.tiles)-1 {
		return []Tile{pipeMap.tiles[current.y+1][current.x]}, nil
	}
	if yDirection == 1 && current.x < len(pipeMap.tiles[current.y])-1 {
		return []Tile{pipeMap.tiles[current.y][current.x+1]}, nil
	}
	if yDirection == -1 && current.x > 0 {
		return []Tile{pipeMap.tiles[current.y][current.x-1]}, nil
	}
	return []Tile{}, errors.New("No valid left hand tile found")
}

func getRightHandTile(pipeMap PipeMap, prev Tile, current Tile) (tiles []Tile, err error) {
	xDirection := current.x - prev.x
	yDirection := current.y - prev.y
	if xDirection == 1 && current.y < len(pipeMap.tiles)-1 {
		return []Tile{pipeMap.tiles[current.y+1][current.x]}, nil
	}
	if xDirection == -1 && current.y > 0 {
		return []Tile{pipeMap.tiles[current.y-1][current.x]}, nil
	}
	if yDirection == 1 && current.x > 0 {
		return []Tile{pipeMap.tiles[current.y][current.x-1]}, nil
	}
	if yDirection == -1 && current.x < len(pipeMap.tiles[current.y])-1 {
		return []Tile{pipeMap.tiles[current.y][current.x+1]}, nil
	}
	return []Tile{}, errors.New("No valid right hand tile found")
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
			newTile := Tile{char, i, lineCount, false, 0}
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

func getNeighbours(pipeMap PipeMap, tile Tile) []Tile {
	neighbours := []Tile{}
	if tile.x > 0 && !pipeMap.tiles[tile.y][tile.x-1].hasBeenVisited {
		neighbours = append(neighbours, pipeMap.tiles[tile.y][tile.x-1])
	}
	if tile.x < len(pipeMap.tiles[tile.y])-1 && !pipeMap.tiles[tile.y][tile.x+1].hasBeenVisited {
		neighbours = append(neighbours, pipeMap.tiles[tile.y][tile.x+1])
	}
	if tile.y > 0 && !pipeMap.tiles[tile.y-1][tile.x].hasBeenVisited {
		neighbours = append(neighbours, pipeMap.tiles[tile.y-1][tile.x])
	}
	if tile.y < len(pipeMap.tiles)-1 && !pipeMap.tiles[tile.y+1][tile.x].hasBeenVisited {
		neighbours = append(neighbours, pipeMap.tiles[tile.y+1][tile.x])
	}
	return neighbours
}

func floodFill(pipeMap *PipeMap, loop []Tile, start Tile, color int) {
	result := []Tile{}
	result = append(result, start)
	pipeMap.tiles[start.y][start.x].group = color
	for len(result) > 0 {
		neigh := getNeighbours(*pipeMap, result[0])
		for _, n := range neigh {
			if slices.IndexFunc(loop, func(t Tile) bool { return hasSamePosition(t, n) }) == -1 && n.group != color {
				result = append(result, n)
				pipeMap.tiles[n.y][n.x].group = color
			}
		}
		result = result[1:]
	}
}

func hasSamePosition(a Tile, b Tile) bool {
	return a.x == b.x && a.y == b.y
}

func main() {
	result := parseData("../data.txt")
	// result := parseData("../sample.txt")
	previousTile := result.loop[0]
	rightHandArray := []Tile{}
	leftHandArray := []Tile{}
	for i := 1; i < len(result.loop); i++ {
		listRight, err := getRightHandTile(result, previousTile, result.loop[i])
		if err == nil {
			for _, tile := range listRight {
				if !tile.hasBeenVisited {
					rightHandArray = append(rightHandArray, tile)
				}
			}
		}
		listLeft, err := getLeftHandTile(result, previousTile, result.loop[i])
		if err == nil {
			for _, tile := range listLeft {
				if !tile.hasBeenVisited {
					leftHandArray = append(leftHandArray, tile)
				}
			}
		}
		previousTile = result.loop[i]
	}

	for _, tile := range leftHandArray {
		floodFill(&result, result.loop, tile, 2)
	}
	for _, tile := range rightHandArray {
		floodFill(&result, result.loop, tile, 1)
	}

	grp1Count := 0
	grp2Count := 0
	for _, line := range result.tiles {
		for _, tile := range line {
			if tile.group == 1 {
				grp1Count++
			}
			if tile.group == 2 {
				grp2Count++
			}
		}
	}
	fmt.Println(result)
	fmt.Println(grp1Count)
	fmt.Println(grp2Count)
}
