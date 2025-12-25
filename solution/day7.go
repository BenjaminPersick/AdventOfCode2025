package solution

import (
	"strconv"
	"strings"
)

// Enum for the type of manifold diagram space.
type ManifoldSpace int

const (
	Empty    ManifoldSpace = 0  // An empty space.
	Splitter ManifoldSpace = -1 // A space occupied by a beam splitter
	Tachyon  ManifoldSpace = 1  // A space occupied by a tachyon beam (or single particle)
	Source   ManifoldSpace = -2 // The space from which the initial beam emanates.
)

// The structured form of the day 7 puzzle input.
type Day7 struct {
	grid [][]ManifoldSpace
}

// Translates the raw text input for day 7 into its structured form.
func (d *Day7) Init(input string) {
	input = strings.TrimSpace(input)

	rows := strings.Split(input, "\n")
	numRows := len(rows)
	numCols := len(rows[0])

	var spaceType = map[rune]ManifoldSpace{
		'.': Empty,
		'^': Splitter,
		'S': Source,
	}

	d.grid = make([][]ManifoldSpace, numRows)
	for i := range numRows {
		d.grid[i] = make([]ManifoldSpace, numCols)
		for j := range numCols {
			d.grid[i][j] = spaceType[rune(rows[i][j])]
		}
	}
}

// Returns the solution for part one of day 7.
func (d Day7) PartOne() string {
	splitCount := 0

	numRows := len(d.grid)
	numCols := len(d.grid[0])

	// create the initial beam in the second row underneath the source
	var i, j int
	for i = range numCols {
		if d.grid[0][i] == Source {
			d.grid[1][i] = Tachyon
			break
		}
	}

	// literally just simulate the beam descending and being split
	for i = 1; i < numRows-1; i++ {
		for j = range numCols {
			if d.grid[i][j] == Tachyon {
				if d.grid[i+1][j] == Splitter {
					splitCount++
					d.grid[i+1][j-1] = Tachyon
					d.grid[i+1][j+1] = Tachyon
				} else {
					d.grid[i+1][j] = Tachyon
				}
			}
		}
	}

	return strconv.Itoa(splitCount)
}

// Returns the solution for part two of day 7.
func (d Day7) PartTwo() string {
	numRows := len(d.grid)
	numCols := len(d.grid[0])

	// (we don't need to create the initial particle here bcause that was already done by PartOne)s

	/*for the many-worlds model, we can just represent a tachyon space with the number of particles occupying
	that space across all timelines*/
	var i, j int
	var currSpaceVal ManifoldSpace
	for i = 1; i < numRows-1; i++ {
		// clear the modifications that were made to the next row by PartOne
		for j = range numCols {
			if d.grid[i+1][j] == Tachyon {
				d.grid[i+1][j] = Empty
			}
		}

		// calculate the values for the next row
		for j = range numCols {
			currSpaceVal = d.grid[i][j]

			if currSpaceVal >= Tachyon {
				if d.grid[i+1][j] == Splitter {
					d.grid[i+1][j-1] += currSpaceVal
					d.grid[i+1][j+1] += currSpaceVal
				} else {
					d.grid[i+1][j] += currSpaceVal
				}
			}
		}
	}

	// calculate the number of timelines by adding up all of the values in the last row
	numTimelines := 0
	for i = range numCols {
		numTimelines += int(d.grid[numRows-1][i])
	}

	return strconv.Itoa(numTimelines)
}
