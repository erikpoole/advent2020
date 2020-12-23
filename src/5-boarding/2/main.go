package main

import (
	"fmt"
	"log"

	"../../utils"
)

const (
	xshift = 3
	yShift = 1
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	boardingPasses := utils.GetLines(file)

	seats := [128][8]int{}
	for _, pass := range boardingPasses {
		row := decodeSeat(pass[:len(pass)-3], 0, 127)
		column := decodeSeat(pass[len(pass)-3:], 0, 7)
		seats[row][column] = 1
	}

	row, column, ok := findSeat(seats)
	if !ok {
		log.Fatal("seat not found")
	}

	fmt.Println((row * 8) + column)
}

func decodeSeat(boardingPass string, lower, upper int) int {
	if upper == lower {
		return upper
	}

	if len(boardingPass) == 0 {
		log.Fatal("boarding pass could not determine a seat")
	}

	positionByte := boardingPass[0]
	middle := lower + (upper-lower)/2
	if positionByte == 'F' || positionByte == 'L' {
		upper = middle
	} else if positionByte == 'B' || positionByte == 'R' {
		lower = middle + 1
	} else {
		log.Fatal(fmt.Sprintf("invalid input: %v", positionByte))
	}

	return decodeSeat(boardingPass[1:], lower, upper)
}

func findSeat(seats [128][8]int) (int, int, bool) {
	foundExistingSeat := false
	for rowIndex := range seats {
		for colIndex := range seats[rowIndex] {
			if !foundExistingSeat {
				if seats[rowIndex][colIndex] == 1 {
					foundExistingSeat = true
				}
				continue
			}

			if seats[rowIndex][colIndex] == 0 {
				return rowIndex, colIndex, true
			}
		}
	}
	return 0, 0, false
}
