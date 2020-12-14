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

	splitIndexes := strings.Split(limits, "-")

	if len(splitIndexes) != 2 {
		return false, fmt.Errorf("invalid indexes split length")
	}

	index1, err := strconv.Atoi(splitIndexes[0])
	if err != nil {
		return false, err
	}

	index2, err := strconv.Atoi(splitIndexes[1])
	if err != nil {
		return false, err
	}

	requiredChar := splitPolicy[1][:1]
	password := splitPolicy[2]

	matchesIndex1 := doesMatchIndexValue(password, []byte(requiredChar)[0], index1-1)
	matchesIndex2 := doesMatchIndexValue(password, []byte(requiredChar)[0], index2-1)

	if (matchesIndex1 || matchesIndex2) && !(matchesIndex1 && matchesIndex2) {
		return true, nil
	}

	return false, nil
}

func doesMatchIndexValue(inputString string, inputByte byte, index int) bool {
	if inputString[index] == inputByte {
		return true
	}
	return false
}
