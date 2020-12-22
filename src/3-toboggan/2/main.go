package main

import (
	"fmt"

	"../../utils"
)

type slope struct {
	xShift int
	yShift int
}

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	treeRows := utils.GetLines(file)
	treeGrid := stringListToRuneGrid(treeRows)

	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, workingSlope := range slopes {
		product *= getTreesHit(treeGrid, workingSlope.xShift, workingSlope.yShift)
	}

	fmt.Println(product)
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

func getTreesHit(treeGrid [][]rune, xShift, yShift int) int {
	xCurr := 0
	yCurr := 0
	treesHit := 0
	for yCurr < len(treeGrid) {
		if treeGrid[yCurr][xCurr] == '#' {
			treesHit++
		}
		xCurr = (xCurr + xShift) % len(treeGrid[0])
		yCurr = yCurr + yShift
	}

	return treesHit
}
