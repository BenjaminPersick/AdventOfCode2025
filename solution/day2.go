package solution

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Range struct {
	low  int
	high int
}

type Day2 struct {
	ranges []Range
}

func (d *Day2) Init(input string) {
	input = strings.TrimSpace(input)
	rangeStrings := strings.Split(input, ",")

	numRanges := len(rangeStrings)

	d.ranges = make([]Range, numRanges)
	var currRangeStr string
	var currDashIndex, currLow, currHigh int
	var err error
	for i := range numRanges {
		currRangeStr = rangeStrings[i]
		currDashIndex = strings.IndexRune(currRangeStr, '-')

		if currDashIndex < 0 {
			log.Fatal("Invalid range string - no dash found.")
		}

		currLow, err = strconv.Atoi(currRangeStr[:currDashIndex])

		if err != nil {
			log.Fatal(err)
		}

		currHigh, err = strconv.Atoi(currRangeStr[currDashIndex+1:])

		if err != nil {
			log.Fatal(err)
		}

		d.ranges[i] = Range{currLow, currHigh}
	}
}

func (d Day2) PartOne() string {
	invalidIDSum := 0

	var currRange Range
	var j int
	for i := range len(d.ranges) {
		currRange = d.ranges[i]

		for j = currRange.low; j <= currRange.high; j++ {
			if IsDoubleSequence(j) {
				invalidIDSum += j
			}
		}
	}

	return fmt.Sprintf("%d", invalidIDSum)
}

func (d Day2) PartTwo() string {
	invalidIDSum := 0

	var currRange Range
	var j int
	for i := range len(d.ranges) {
		currRange = d.ranges[i]

		for j = currRange.low; j <= currRange.high; j++ {
			if IsRepeatedSequence(j) {
				invalidIDSum += j
			}
		}
	}

	return fmt.Sprintf("%d", invalidIDSum)
}

func IsDoubleSequence(num int) bool {
	// get number of digits
	numDigits := 0
	tempNum := num
	for tempNum != 0 {
		tempNum /= 10
		numDigits++
	}

	if numDigits%2 != 0 {
		return false
	}

	// get upper and lower halfs of the number
	halfPowerOfTen := int(math.Pow10(numDigits / 2))
	upperHalf := num / halfPowerOfTen
	lowerHalf := num - (upperHalf * halfPowerOfTen)

	return upperHalf == lowerHalf
}

func IsRepeatedSequence(num int) bool {
	numString := strconv.Itoa(num)
	numDigits := len(numString)

	// test all possible sequence lengths
	var potentialSeq string
	var startIndex, endIndex int
	var mismatchFound bool
	for seqSize := 1; seqSize <= numDigits/2; seqSize++ {
		// skip sequence lengths by which the number of digits is not divisible
		if numDigits%seqSize != 0 {
			continue
		}

		potentialSeq = numString[:seqSize]

		startIndex = seqSize
		endIndex = startIndex + seqSize
		mismatchFound = false
		for endIndex <= numDigits {
			if numString[startIndex:endIndex] != potentialSeq {
				mismatchFound = true
				break
			}

			startIndex += seqSize
			endIndex += seqSize
		}

		if !mismatchFound {
			return true
		}
	}

	return false
}
