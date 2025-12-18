package solution

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DialMove struct {
	isLeft   bool
	turnDist int
}

type Day1 struct {
	moves []DialMove
}

func (d *Day1) Init(input string) {
	lines := strings.Split(input, "\n")

	numLines := len(lines)
	d.moves = make([]DialMove, numLines)

	var currLine string
	for i := 0; i < numLines; i++ {
		currLine = lines[i]

		if len(currLine) < 1 {
			break
		}

		leftRotation := rune(currLine[0]) == 'L'
		distance, err := strconv.Atoi(currLine[1:])

		if err != nil {
			log.Fatal(err)
		}

		d.moves[i] = DialMove{leftRotation, distance}
	}
}

func (d Day1) PartOne() string {
	zeroCount := 0
	currPosition := 50

	var currMove DialMove
	for i := 0; i < len(d.moves); i++ {
		currMove = d.moves[i]

		if currMove.isLeft {
			currPosition -= currMove.turnDist
		} else {
			currPosition += currMove.turnDist
		}

		currPosition %= 100

		if currPosition < 0 {
			currPosition = 100 + currPosition
		}

		if currPosition == 0 {
			zeroCount++
		}
	}

	return fmt.Sprintf("%d", zeroCount)
}

func (d Day1) PartTwo() string {
	zeroCount := 0
	currPosition := 50

	var currMove DialMove
	var initialPosition int
	for i := 0; i < len(d.moves); i++ {
		currMove = d.moves[i]

		if currMove.isLeft {
			initialPosition = currPosition
			currPosition -= currMove.turnDist

			if currPosition <= 0 {
				if initialPosition != 0 {
					zeroCount++
				}

				zeroCount += (currPosition / -100)
			}
		} else {
			currPosition += currMove.turnDist

			zeroCount += currPosition / 100
		}

		currPosition %= 100

		if currPosition < 0 {
			currPosition = 100 + currPosition
		}
	}

	return fmt.Sprintf("%d", zeroCount)
}
