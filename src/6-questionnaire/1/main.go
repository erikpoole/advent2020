package main

import (
	"fmt"

	"../../utils"
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	responses := utils.GetLines(file)

	groups := []map[rune]struct{}{}
	currentGroup := map[rune]struct{}{}
	for _, response := range responses {
		if response == "" {
			groups = append(groups, currentGroup)
			currentGroup = map[rune]struct{}{}
			continue
		}
		for _, affirmativeAnswers := range response {
			currentGroup[affirmativeAnswers] = struct{}{}
		}
	}

	groups = append(groups, currentGroup)

	count := 0
	for _, group := range groups {
		count += len(group)
	}

	fmt.Println(count)
}
