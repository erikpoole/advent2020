package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"../../utils"
)

func main() {
	file := utils.GetInputFile()
	defer file.Close()

	expenses := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		expenses = append(expenses, expense)
	}

	values2020 := findThreeValuesSumming2020(expenses)

	fmt.Println(values2020[0])
	fmt.Println(values2020[1])
	fmt.Println(values2020[2])

	total := 1
	for _, value := range values2020 {
		total *= value
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findThreeValuesSumming2020(expenses []int) []int {
	for _, value1 := range expenses {
		for _, value2 := range expenses {
			if value1+value2 > 2020 {
				continue
			}
			for _, value3 := range expenses {
				if value1+value2+value3 == 2020 {
					return []int{value1, value2, value3}
				}
			}
		}
	}
	return []int{}
}
