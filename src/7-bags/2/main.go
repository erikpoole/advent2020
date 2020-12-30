package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../utils"
)

type numberAndColor struct {
	number int
	color  string
}

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	baglines := utils.GetLines(file)

	bagsToChildren := createBagsToChildren(baglines)

	fmt.Println(findBagsInBag("shiny gold", bagsToChildren))
}

func createBagsToChildren(bagLines []string) map[string][]numberAndColor {
	bagsToChildren := map[string][]numberAndColor{}
	for _, bagline := range bagLines {
		parentBagColor, childBagColors := parseBagLine(bagline)
		bagsToChildren[parentBagColor] = childBagColors
	}

	return bagsToChildren
}

func parseBagLine(bagLine string) (string, []numberAndColor) {
	splitBagLine := strings.Split(bagLine, " contain ")
	if len(splitBagLine) != 2 {
		log.Fatal(fmt.Errorf("invalid bag split during parsing: %v", bagLine))
	}

	parentBagColor := strings.Join(strings.Split(splitBagLine[0], " ")[:2], " ")
	childBagsLine := splitBagLine[1]

	if strings.Contains(childBagsLine, "no other bags") {
		return parentBagColor, []numberAndColor{}
	}

	childBags := strings.Split(childBagsLine, ", ")
	childBagColors := []numberAndColor{}
	for _, childBag := range childBags {
		splitChildBags := strings.Split(childBag, " ")
		if len(splitChildBags) != 4 {
			log.Fatal(fmt.Errorf("invalid bag split during parsing: %v", bagLine))
		}
		childBagNumber, err := strconv.Atoi(splitChildBags[0])
		if err != nil {
			log.Fatal(fmt.Errorf("invalid bag number during parsing: %v", bagLine))
		}
		childBagColor := strings.Join(splitChildBags[1:3], " ")
		childBagColors = append(childBagColors, numberAndColor{childBagNumber, childBagColor})
	}

	return parentBagColor, childBagColors
}

func findBagsInBag(bagColor string, bagsToChildren map[string][]numberAndColor) int {
	childBags := bagsToChildren[bagColor]
	totalBags := 0
	for _, childBag := range childBags {
		totalBags += childBag.number * (1 + findBagsInBag(childBag.color, bagsToChildren))
	}

	return totalBags
}
