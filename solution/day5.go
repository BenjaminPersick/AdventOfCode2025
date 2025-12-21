package solution

import (
	"log"
	"strconv"
	"strings"
)

type Day5 struct {
	freshRanges  []Range
	availableIDs []int
}

func (d *Day5) Init(input string) {
	input = strings.TrimSpace(input)

	// split data into ranges and IDs
	database := strings.Split(input, "\n\n")
	rangeStrings := strings.Split(database[0], "\n")
	idStrings := strings.Split(database[1], "\n")

	numRanges := len(rangeStrings)
	numIDs := len(idStrings)

	var i, dashIndex, currMin, currMax, currID int
	var currRangeStr string
	var err error

	// get ranges
	d.freshRanges = make([]Range, numRanges)
	for i = range numRanges {
		currRangeStr = rangeStrings[i]
		dashIndex = strings.IndexRune(currRangeStr, '-')

		currMin, err = strconv.Atoi(currRangeStr[:dashIndex])

		if err != nil {
			log.Fatal(err)
		}

		currMax, err = strconv.Atoi(currRangeStr[dashIndex+1:])

		if err != nil {
			log.Fatal(err)
		}

		d.freshRanges[i] = Range{currMin, currMax}
	}

	// get IDs
	d.availableIDs = make([]int, numIDs)
	for i = range numIDs {
		currID, err = strconv.Atoi(idStrings[i])

		if err != nil {
			log.Fatal(err)
		}

		d.availableIDs[i] = currID
	}
}

func (d Day5) PartOne() string {
	freshCount := 0

	var j int
	for i := range len(d.availableIDs) {
		for j = range d.freshRanges {
			if InRange(d.availableIDs[i], d.freshRanges[j]) {
				freshCount++
				break
			}
		}
	}

	return strconv.Itoa(freshCount)
}

func (d Day5) PartTwo() string {
	// merge overlapping ranges together before computing answer
	mergedRanges := make([]Range, 0)

	var i, j, upperOverlappingRangeIndex int
	var currRange Range
	var currMergedRange, lowerOverlappingRange, upperOverlappingRange *Range
	var shouldAppend bool
	for i = range len(d.freshRanges) {
		currRange = d.freshRanges[i]

		shouldAppend = true

		// check whether the current range we are merging in overlapps with any of the current merged ranges
		lowerOverlappingRange = nil
		upperOverlappingRange = nil
		upperOverlappingRangeIndex = -1
		for j = range len(mergedRanges) {
			currMergedRange = &mergedRanges[j]

			if InRange(currRange.low, *currMergedRange) {
				shouldAppend = false

				// we don't need to do any merging if the current range is completely within another range
				if InRange(currRange.high, *currMergedRange) {
					break
				}

				lowerOverlappingRange = currMergedRange
				continue
			}

			if InRange(currRange.high, *currMergedRange) {
				shouldAppend = false
				upperOverlappingRange = currMergedRange
				upperOverlappingRangeIndex = j

				continue
			}

			// handle the case where an existing range in mergedRanges wis completely encapsulated by currRange
			if InRange(currMergedRange.low, currRange) && InRange(currMergedRange.high, currRange) {
				// overwrite the values of currMergedRange with those of currRange
				currMergedRange.low = currRange.low
				currMergedRange.high = currRange.high

				shouldAppend = false
			}
		}

		if lowerOverlappingRange != nil && upperOverlappingRange != nil {
			// current range overlaps with two ranges in the mergedRanges array => merge all 3 together
			lowerOverlappingRange.high = upperOverlappingRange.high

			// delete upper range from mergedRange after the merge by overwriting it with the last element in the array and then curring off the last element
			mergedRanges[upperOverlappingRangeIndex], mergedRanges = mergedRanges[len(mergedRanges)-1], mergedRanges[:len(mergedRanges)-1]
		} else if lowerOverlappingRange != nil {
			// current range's lower end overlaps with a range in mergedRanges
			lowerOverlappingRange.high = currRange.high
		} else if upperOverlappingRange != nil {
			// current range's upper end overlapps with a range in mergedRanges
			upperOverlappingRange.low = currRange.low
		}

		if shouldAppend {
			mergedRanges = append(mergedRanges, currRange)
		}
	}

	// get the total number of fresh IDs from the merged ranges
	totalFreshCount := 0
	for i = range len(mergedRanges) {
		currRange = mergedRanges[i]

		totalFreshCount += (currRange.high - currRange.low) + 1
	}

	return strconv.Itoa(totalFreshCount)
}

// Determines whether the given value is within the given (inclusive) range
func InRange(val int, rng Range) bool {
	return val >= rng.low && val <= rng.high
}
