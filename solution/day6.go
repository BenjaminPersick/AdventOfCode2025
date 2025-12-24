package solution

import (
	"log"
	"strconv"
	"strings"
)

// A column from the math homework sheet.
type MathColumn struct {
	operands []int // The operands of this column of the math homework sheet.
	isMult   bool  // Whether this column is a multiplication column or an addition column.
}

// The structured form of the day 6 puzzle input.
type Day6 struct {
	columns           []MathColumn // All of the columns from the math homework sheet.
	cephalopodColumns []MathColumn // All of the columns from the math homework sheet, but read according to cephalopod digit arrangement spec.
}

// Translates the raw text input for day 6 into its structured form.
func (d *Day6) Init(input string) {
	input = strings.TrimSuffix(input, "\n")

	rows := strings.Split(input, "\n")
	numRows := len(rows)
	splitRows := make([][]string, numRows)

	var i int
	for i = range numRows {
		splitRows[i] = strings.Fields(rows[i])
	}

	numCols := len(splitRows[0])

	var currOperands, currCephOperands []int
	var j, k, operand, endIndex, numCephOperands, currDigitIndex int
	var err error
	var currIsMult bool
	var operandStr string
	var currRune rune
	numHumanOperands := numRows - 1
	startIndex := 0
	operatorRow := rows[numRows-1]

	d.columns = make([]MathColumn, numCols)
	d.cephalopodColumns = make([]MathColumn, numCols)
	for i = range numCols {
		// build column based off of human notation
		currOperands = make([]int, numHumanOperands)

		for j = range numHumanOperands {
			operand, err = strconv.Atoi(splitRows[j][i])

			if err != nil {
				log.Fatal(err)
			}

			currOperands[j] = operand
		}

		currIsMult = splitRows[numRows-1][i] == "*"

		d.columns[i] = MathColumn{currOperands, currIsMult}

		// build column based off of cephalopod notation

		// use the row containing the operators to determine the range of indices corresponding to the current column
		endIndex = startIndex + 1
		for (endIndex < len(operatorRow)-1 && rune(operatorRow[endIndex+1]) == ' ') || endIndex == len(operatorRow)-1 {
			endIndex++
		}

		numCephOperands = endIndex - startIndex

		currCephOperands = make([]int, numCephOperands)
		for j = range numCephOperands {
			currDigitIndex = startIndex + j
			operandStr = ""

			for k = range numRows - 1 {
				currRune = rune(rows[k][currDigitIndex])

				if currRune != ' ' {
					operandStr += string(currRune)
				}
			}

			operand, err = strconv.Atoi(operandStr)

			if err != nil {
				log.Fatal(err)
			}

			currCephOperands[j] = operand
		}

		d.cephalopodColumns[i] = MathColumn{currCephOperands, currIsMult}

		startIndex = endIndex + 1
	}
}

// Returns the solution for part one of day 6.
func (d Day6) PartOne() string {

	return strconv.Itoa(SumColumnSolutions(d.columns))
}

// Returns the solution for part two of day 6.
func (d Day6) PartTwo() string {
	return strconv.Itoa(SumColumnSolutions(d.cephalopodColumns))
}

// Calculates the sum of the answers of all MathColumns in the given array.
func SumColumnSolutions(cols []MathColumn) int {
	grandTotal := 0

	var currCol MathColumn
	var currSolution, j int
	for i := range len(cols) {
		currCol = cols[i]
		currSolution = currCol.operands[0]

		for j = 1; j < len(currCol.operands); j++ {
			if currCol.isMult {
				currSolution *= currCol.operands[j]
			} else {
				currSolution += currCol.operands[j]
			}
		}

		grandTotal += currSolution
	}

	return grandTotal
}
