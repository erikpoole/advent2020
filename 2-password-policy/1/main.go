package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../utils"
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	policies := utils.GetLines(file)

	var validCount int
	for _, policy := range policies {
		valid, err := isPasswordPolicyValid(policy)
		if err != nil {
			log.Fatal(fmt.Errorf("Invalid line input - %v: %w", policy, err))
		}

		if valid {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func isPasswordPolicyValid(policy string) (bool, error) {
	splitPolicy := strings.Split(policy, " ")

	if len(splitPolicy) != 3 {
		return false, fmt.Errorf("invalid policy split length")
	}

	limits := splitPolicy[0]

	splitLimits := strings.Split(limits, "-")

	if len(splitLimits) != 2 {
		return false, fmt.Errorf("invalid limits split length")
	}

	minAllowed, err := strconv.Atoi(splitLimits[0])
	if err != nil {
		return false, err
	}

	maxAllowed, err := strconv.Atoi(splitLimits[1])
	if err != nil {
		return false, err
	}

	requiredChar := splitPolicy[1][:1]
	password := splitPolicy[2]

	occurences := strings.Count(password, requiredChar)

	if occurences <= maxAllowed && occurences >= minAllowed {
		return true, nil
	}

	return false, nil
}
