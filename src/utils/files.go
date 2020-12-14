package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetInputFile() *os.File {
	if len(os.Args) != 2 {
		log.Fatal(fmt.Errorf("Expected arg - inputfile path"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func GetLines(file *os.File) []string {
	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
