package solution

import (
	"strconv"
	"strings"
)

type Day4 struct {
	grid [][]bool
}

func (d *Day4) Init(input string) {
	input = strings.TrimSpace(input)

	rows := strings.Split(input, "\n")
	numRows := len(rows)

	d.grid = make([][]bool, numRows)

	var j, rowSize int
	var currRow string
	for i := range numRows {
		currRow = rows[i]
		rowSize = len(currRow)

		d.grid[i] = make([]bool, rowSize)
		for j = range rowSize {
			d.grid[i][j] = rune(currRow[j]) == '@'
		}
	}
}

func (d Day4) PartOne() string {
	accessibleRollCount := 0

	var j int
	for i := range len(d.grid) {
		for j = range len(d.grid[i]) {
			if d.grid[i][j] && CountAdjacentRolls(d.grid, i, j) < 4 {
				accessibleRollCount++
			}
		}
	}

	return strconv.Itoa(accessibleRollCount)
}

func (d Day4) PartTwo() string {
	removalCount := 0

	var i, j int

	rollsRemoved := true
	for rollsRemoved {
		rollsRemoved = false

		for i = range len(d.grid) {
			for j = range len(d.grid[i]) {
				if d.grid[i][j] && CountAdjacentRolls(d.grid, i, j) < 4 {
					d.grid[i][j] = false
					rollsRemoved = true
					removalCount++
				}
			}
		}
	}

	return strconv.Itoa(removalCount)
}

func CountAdjacentRolls(grid [][]bool, row int, col int) int {
	indicesToTest := [8][2]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}

	rowLength := len(grid[0])
	colHeight := len(grid)

	rollCount := 0

	var currRow, currCol int
	for i := range len(indicesToTest) {
		currRow = indicesToTest[i][0]
		currCol = indicesToTest[i][1]

		if currRow < 0 || currRow >= colHeight || currCol < 0 || currCol >= rowLength {
			continue
		}

		if grid[currRow][currCol] {
			rollCount++
		}
	}

	return rollCount
}
