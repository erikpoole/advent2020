package main

import (
	"container/list"
	"fmt"
	"log"
	"strings"

	"../../utils"
)

// keep map of bags and their parent bags
// key: color, value: array of parent bags color

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	baglines := utils.GetLines(file)

	bagsToParents := createBagsToParents(baglines)

	// start with gold bag, add parents to queue
	// for each value in queue, add parents to queue
	// if no parents, add to set of valid outer bags
	// return count of valid outer bags
	outerBags := map[string]struct{}{}

	queue := list.New()
	queue.PushBack("shiny gold")

	for queue.Len() != 0 {
		currentBag := queue.Front()
		queue.Remove(currentBag)
		bagColor := currentBag.Value.(string)
		outerBags[bagColor] = struct{}{}
		bagParents, _ := bagsToParents[bagColor]

		for _, bagParent := range bagParents {
			queue.PushBack(bagParent)
		}
	}

	// implementation will contain "shiny gold", which we want to omit
	fmt.Println(len(outerBags) - 1)
}

// map construction - loop through each line
// for each child bag -
// if not child bag key exists -
// add child bag key
// add parent bag to child key array

func createBagsToParents(bagLines []string) map[string][]string {
	bagsToParents := map[string][]string{}
	for _, bagline := range bagLines {
		parentBagColor, childBagColors := parseBagLine(bagline)
		for _, childBagColor := range childBagColors {
			parentBagColors, ok := bagsToParents[childBagColor]
			if !ok {
				parentBagColors = []string{}
				bagsToParents[childBagColor] = parentBagColors
			}
			bagsToParents[childBagColor] = append(parentBagColors, parentBagColor)
		}
	}

	return bagsToParents
}

func parseBagLine(bagLine string) (string, []string) {
	splitBagLine := strings.Split(bagLine, " contain ")
	if len(splitBagLine) != 2 {
		log.Fatal(fmt.Errorf("invalid bag split during parsing: %v", bagLine))
	}

	parentBagColor := strings.Join(strings.Split(splitBagLine[0], " ")[:2], " ")
	childBagsLine := splitBagLine[1]

	if strings.Contains(childBagsLine, "no other bags") {
		return parentBagColor, []string{}
	}

	childBags := strings.Split(childBagsLine, ", ")
	childBagColors := []string{}
	for _, childBag := range childBags {
		splitChildBags := strings.Split(childBag, " ")
		if len(splitChildBags) != 4 {
			log.Fatal(fmt.Errorf("invalid bag split during parsing: %v", bagLine))
		}
		childBagColor := strings.Join(splitChildBags[1:3], " ")
		childBagColors = append(childBagColors, childBagColor)
	}

	return parentBagColor, childBagColors
}
