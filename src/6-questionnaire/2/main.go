package main

import (
	"fmt"

	"../../utils"
)

const (
	minRuneValue = 'a'
	maxRuneValue = 'z'
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	responses := utils.GetLines(file)

	count := 0
	currentGroup := []map[rune]struct{}{}
	for _, response := range responses {
		if response == "" {
			// compare currentGroup
			commonalities := findCommonalities(currentGroup, minRuneValue, maxRuneValue)
			count += len(commonalities)
			currentGroup = []map[rune]struct{}{}
			continue
		}

		affirmativeAnswers := map[rune]struct{}{}
		for _, affirmativeAnswer := range response {
			affirmativeAnswers[affirmativeAnswer] = struct{}{}
		}
		currentGroup = append(currentGroup, affirmativeAnswers)
	}

	// compare last group
	commonalities := findCommonalities(currentGroup, minRuneValue, maxRuneValue)
	count += len(commonalities)

	fmt.Println(count)
}

func findCommonalities(maps []map[rune]struct{}, lowerBound, upperBound rune) string {
	commonalities := ""
	for i := lowerBound; i <= upperBound; i++ {
		validInAll := true
		for _, individualMap := range maps {
			if _, ok := individualMap[i]; !ok {
				validInAll = false
			}
		}
		if validInAll {
			commonalities += string(i)
		}
	}

	return commonalities
}
