package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BenjaminPersick/AdventOfCode2025/solution"
)

func main() {
	days := []solution.SolutionPair{
		&solution.Day1{},
	}

	var currPair solution.SolutionPair
	for i := 0; i < len(days); i++ {
		currPair = days[i]

		currPair.Init(readFileContent(fmt.Sprintf("./input/%d.txt", i+1)))
		fmt.Printf("%s\n", currPair.PartOne())
		fmt.Printf("%s\n", currPair.PartTwo())
	}
}

func readFileContent(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
