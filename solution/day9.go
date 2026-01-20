package solution

import (
	"log"
	"strconv"
	"strings"
)

// Represents a point in 2D space
type Vec2 struct {
	x int
	y int
}

// The structured form of the day 9 puzzle input.
type Day9 struct {
	tileCoords []Vec2 // The list of all red tile coordinates.
}

// Translates the raw text input for day 9 into its structured form.
func (d *Day9) Init(input string) {
	input = strings.TrimSpace(input)

	rows := strings.Split(input, "\n")
	numRows := len(rows)

	d.tileCoords = make([]Vec2, numRows)

	var currX, currY, currCommaIndex int
	var err error
	var currRow string
	for i := range numRows {
		currRow = rows[i]
		currCommaIndex = strings.IndexRune(currRow, ',')

		currX, err = strconv.Atoi(currRow[:currCommaIndex])

		if err != nil {
			log.Fatal(err)
		}

		currY, err = strconv.Atoi(currRow[currCommaIndex+1:])

		if err != nil {
			log.Fatal(err)
		}

		d.tileCoords[i] = Vec2{currX, currY}
	}
}

// Returns the solution for part one of day 9.
func (d Day9) PartOne() string {
	greatestArea := 0

	numCoords := len(d.tileCoords)
	var i, j, currArea int
	for i = range numCoords - 1 {
		for j = i + 1; j < numCoords; j++ {
			currArea = TileArea(d.tileCoords[i], d.tileCoords[j])

			if currArea > greatestArea {
				greatestArea = currArea
			}
		}
	}

	return strconv.Itoa(greatestArea)
}

// Returns the solution for part two of day 9.
func (d Day9) PartTwo() string {
	return "not implemented"
}

// Gets the area of the rectangle formed by the two tiles with the given coordinates.
func TileArea(corner1, corner2 Vec2) int {
	dx := corner2.x - corner1.x
	dy := corner2.y - corner1.y

	if dx < 0 {
		dx *= -1
	}

	if dy < 0 {
		dy *= -1
	}

	dx++
	dy++

	return dx * dy
}
