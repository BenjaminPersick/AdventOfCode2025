package solution

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Represents a pair of junction boxes and the distance between them.
type BoxPair struct {
	box1Index int
	box2Index int
	distance  float64
}

// Represents a point in 3D space.
type Vec3 struct {
	x int
	y int
	z int
}

// The structured form of the day 8 puzzle input.
type Day8 struct {
	junctionBoxes []Vec3    // The list of all junction box coordinates.
	boxPairs      []BoxPair // All possible pairings of boxes in ascending order of distance between boxes.
}

// Translates the raw text input for day 8 into its structured form.
func (d *Day8) Init(input string) {
	// get list of all junction boxes
	input = strings.TrimSpace(input)

	rows := strings.Split(input, "\n")
	numRows := len(rows)

	d.junctionBoxes = make([]Vec3, numRows)

	currCoords := make([]int, 3)
	var currCoordStrings []string
	var err error
	for i := range numRows {
		currCoordStrings = strings.Split(rows[i], ",")

		for j := range 3 {
			currCoords[j], err = strconv.Atoi(currCoordStrings[j])

			if err != nil {
				log.Fatal(err)
			}
		}

		d.junctionBoxes[i] = Vec3{currCoords[0], currCoords[1], currCoords[2]}
	}

	numBoxes := len(d.junctionBoxes)

	// https://math.stackexchange.com/questions/4038934/formula-for-number-of-pairs-that-can-be-made-from-n-items
	numPairs := (numBoxes * (numBoxes - 1)) / 2

	// get the distances between all possible pairs of junction boxes
	d.boxPairs = make([]BoxPair, numPairs)[:]

	var i, j int
	pairIndex := 0
	for i = 0; i < numBoxes-1; i++ {
		for j = i + 1; j < numBoxes; j++ {
			d.boxPairs[pairIndex] = BoxPair{i, j, Distance(d.junctionBoxes[i], d.junctionBoxes[j])}

			pairIndex++
		}
	}

	// sort the box pairs by distance
	sort.Slice(d.boxPairs, func(i, j int) bool {
		return d.boxPairs[i].distance < d.boxPairs[j].distance
	})
}

// Returns the solution for part one of day 8.
func (d Day8) PartOne() string {
	// set up the array of all of our current circuits and create a map for efficient circuit index lookup
	numBoxes := len(d.junctionBoxes)
	circuits := make([][]int, numBoxes)[:]
	circuitIndexMap := make(map[int]int)

	var i int
	for i = range numBoxes {
		circuits[i] = []int{i}
		circuitIndexMap[i] = i
	}

	// merge the 1000 closest junction boxes
	var currPair BoxPair
	for i = range 1000 {
		currPair = d.boxPairs[i]
		MergeCircuits(&circuits, &circuitIndexMap, currPair.box1Index, currPair.box2Index)
	}

	// sort all circuits by length and multiply the lengths of the top three
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) < len(circuits[j])
	})

	result := 1
	for i = len(circuits) - 1; i >= len(circuits)-3; i-- {
		result *= len(circuits[i])
	}

	return strconv.Itoa(result)
}

// Returns the solution for part two of day 8.
func (d Day8) PartTwo() string {
	// set up the array of all of our current circuits and create a map for efficient circuit index lookup
	numBoxes := len(d.junctionBoxes)
	circuits := make([][]int, numBoxes)[:]
	circuitIndexMap := make(map[int]int)

	var i int
	for i = range numBoxes {
		circuits[i] = []int{i}
		circuitIndexMap[i] = i
	}

	// merge junction boxes until there is one big circuit, and keep track of the previous x-coordinates before each merge
	var currPair BoxPair
	prevX1, prevX2 := 0, 0
	i = 0
	for len(circuits) > 1 {
		currPair = d.boxPairs[i]

		prevX1 = d.junctionBoxes[currPair.box1Index].x
		prevX2 = d.junctionBoxes[currPair.box2Index].x

		MergeCircuits(&circuits, &circuitIndexMap, currPair.box1Index, currPair.box2Index)

		i++
	}

	return strconv.Itoa(prevX1 * prevX2)
}

// Gets the distance between two points in 3D space.
func Distance(p1 Vec3, p2 Vec3) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2) + math.Pow(float64(p2.z-p1.z), 2))
}

// Merges the circuits containg the two given indices.
func MergeCircuits(circuits *[][]int, circuitIndexMap *map[int]int, boxIndex1, boxIndex2 int) {
	// get the indicies of the circuits that the given boxes are in
	circuitIndex1 := (*circuitIndexMap)[boxIndex1]
	circuitIndex2 := (*circuitIndexMap)[boxIndex2]

	// nothing to merge if the boxes are already in the same circuit
	if circuitIndex1 == circuitIndex2 {
		return
	}

	// update the index map before merge
	circuit2 := (*circuits)[circuitIndex2]
	SetCircuitIndexForBoxes(circuitIndexMap, circuit2, circuitIndex1)

	// merge the circuits
	(*circuits)[circuitIndex1] = append((*circuits)[circuitIndex1], circuit2...)

	initialLastCircuitIndex := len(*circuits) - 1
	(*circuits)[circuitIndex2] = (*circuits)[initialLastCircuitIndex]
	(*circuits)[initialLastCircuitIndex] = nil
	(*circuits) = (*circuits)[:initialLastCircuitIndex]

	// update the index map for the circuit that changed indices on account of the deletion at the end of the merge if necessary
	if circuitIndex2 != initialLastCircuitIndex {
		SetCircuitIndexForBoxes(circuitIndexMap, (*circuits)[circuitIndex2], circuitIndex2)
	}
}

// Updates the given circuit box index to circuit index map by mapping all of the given box indices to the given circuit index
func SetCircuitIndexForBoxes(circuitIndexMap *map[int]int, boxIndices []int, circuitIndex int) {
	var currBoxIndex int
	for i := range len(boxIndices) {
		currBoxIndex = boxIndices[i]
		(*circuitIndexMap)[currBoxIndex] = circuitIndex
	}
}
