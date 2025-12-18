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

		dayNum := i + 1

		currPair.Init(readFileContent(fmt.Sprintf("./input/%d.txt", dayNum)))

		fmt.Printf("===================================[Day %d]===================================\n", dayNum)
		fmt.Printf("Part 1: %s\n", currPair.PartOne())
		fmt.Printf("Part 2: %s\n", currPair.PartTwo())
	}
}

func readFileContent(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
