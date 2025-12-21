package solution

import (
	"math"
	"strconv"
	"strings"
)

type Day3 struct {
	batteryBanks [][]int
}

func (d *Day3) Init(input string) {
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")
	numLines := len(lines)

	d.batteryBanks = make([][]int, numLines)

	var currLine string
	var currLineLength int
	var currArray []int
	var currDigitRune rune
	for i := range numLines {
		currLine = lines[i]
		currLineLength = len(currLine)

		currArray = make([]int, currLineLength)
		for j := range currLineLength {
			// convert current character to int
			currDigitRune = rune(currLine[j])
			currArray[j] = int(currDigitRune - '0')
		}

		d.batteryBanks[i] = currArray
	}
}

func (d Day3) PartOne() string {
	var firstDigitIndex, firstDigit, secondDigit, currDigit, currBankSize, j int

	joltageSum := 0

	var currBatteryBank []int
	for i := range len(d.batteryBanks) {
		currBatteryBank = d.batteryBanks[i]
		currBankSize = len(currBatteryBank)

		// get highest digit by comparing all digits except for the last one
		firstDigit = currBatteryBank[0]
		firstDigitIndex = 0
		for j = 1; j < currBankSize-1; j++ {
			currDigit = currBatteryBank[j]

			if currDigit > firstDigit {
				firstDigit = currDigit
				firstDigitIndex = j
			}
		}

		// get highest remaining digit that occurrs after the first
		secondDigit = currBatteryBank[firstDigitIndex+1]
		for j = firstDigitIndex + 2; j < currBankSize; j++ {
			currDigit = currBatteryBank[j]

			if currDigit > secondDigit {
				secondDigit = currDigit
			}
		}

		joltageSum += (10 * firstDigit) + secondDigit
	}

	return strconv.Itoa(joltageSum)
}

func (d Day3) PartTwo() string {
	var currStartIndex, currEndIndex, currDigit, currHighestDigit, currBankSize, j, k int

	joltageSum := 0

	var currBatteryBank []int
	for i := range len(d.batteryBanks) {
		currBatteryBank = d.batteryBanks[i]
		currBankSize = len(currBatteryBank)

		currStartIndex = 0

		// j will be the power of 10 by which we will multiply the current digit
		for j = 11; j >= 0; j-- {
			// get the highest possible current digit while leaving room for the remaining digits
			currHighestDigit = currBatteryBank[currStartIndex]
			currStartIndex++

			currEndIndex = currBankSize - j

			for k = currStartIndex; k < currEndIndex; k++ {
				currDigit = currBatteryBank[k]

				if currDigit > currHighestDigit {
					currHighestDigit = currDigit
					currStartIndex = k + 1
				}
			}

			joltageSum += currHighestDigit * int(math.Pow(10, float64(j)))
		}
	}

	return strconv.Itoa(joltageSum)
}
