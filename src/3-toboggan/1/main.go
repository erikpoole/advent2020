package main

import (
	"fmt"

	"../../utils"
)

const (
	xshift = 3
	yShift = 1
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	treeRows := utils.GetLines(file)
	treeGrid := stringListToRuneGrid(treeRows)

	xCurr := 0
	yCurr := 0
	treesHit := 0
	for yCurr < len(treeGrid) {
		if treeGrid[yCurr][xCurr] == '#' {
			treesHit++
		}
		xCurr = (xCurr + xshift) % len(treeGrid[0])
		yCurr = yCurr + yShift
	}

	fmt.Println(treesHit)
}

func stringListToRuneGrid(list []string) [][]rune {
	var runeGrid [][]rune

	for yIndex, listItem := range list {
		runeGrid = append(runeGrid, []rune{})
		for _, char := range listItem {
			runeGrid[yIndex] = append(runeGrid[yIndex], char)
		}
	}

	return runeGrid
}
