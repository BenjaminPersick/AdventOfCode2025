package solution

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// Represents a point in 2D space
type Vec2 struct {
	x int
	y int
}

// A pair of tiles. Used to represent rectangles or edges.
type TilePair struct {
	start *Vec2
	end   *Vec2
}

// The structured form of the day 9 puzzle input.
type Day9 struct {
	tileCoords []Vec2 // The list of all red tile coordinates.
}

var minX int                                          // The minimum x coordinate of the loop
var horizontalEdges, verticalEdges map[int][]TilePair // Mappings of rows/columns to loop edges that lie on them
var insideLoopCache map[Vec2]bool                     // Cache to prevent TileInsideLoop from re-computing results

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
	// initialize insideLoopCache
	insideLoopCache = make(map[Vec2]bool)

	// get minimum x coordinate and create edge maps
	minX = math.MaxInt

	horizontalEdges = make(map[int][]TilePair)
	verticalEdges = make(map[int][]TilePair)

	numCoords := len(d.tileCoords)
	var currTile, nextTile *Vec2
	var currEdges []TilePair
	var currEdge TilePair
	var exists bool
	var i int
	for i = range numCoords {
		// get current and next tile forming the current edge
		currTile = &d.tileCoords[i]
		if i < numCoords-1 {
			nextTile = &d.tileCoords[i+1]
		} else {
			nextTile = &d.tileCoords[0]
		}

		// update minimum x coord if necessary
		if currTile.x < minX {
			minX = currTile.x
		}

		// add current edge to appropriate map
		currEdge = TilePair{currTile, nextTile}

		if currTile.x == nextTile.x {
			// vertical edge
			currEdges, exists = verticalEdges[currTile.x]

			if exists {
				verticalEdges[currTile.x] = append(currEdges, currEdge)
			} else {
				verticalEdges[currTile.x] = []TilePair{currEdge}
			}
		} else {
			// horizontal edge
			currEdges, exists = horizontalEdges[currTile.y]

			if exists {
				horizontalEdges[currTile.y] = append(currEdges, currEdge)
			} else {
				horizontalEdges[currTile.y] = []TilePair{currEdge}
			}
		}
	}

	// get the largest valid rectangle
	greatestArea := 0

	var j, currArea int
	for i = range numCoords - 1 {
		for j = i + 1; j < numCoords; j++ {
			if RectInsideLoop(TilePair{&d.tileCoords[i], &d.tileCoords[j]}) {
				currArea = TileArea(d.tileCoords[i], d.tileCoords[j])

				if currArea > greatestArea {
					greatestArea = currArea
				}
			}
		}
	}

	return strconv.Itoa(greatestArea)
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

// Determines whether the rectangle with the corners specified by the given TilePair is inside the loop of red/green tiles
func RectInsideLoop(rect TilePair) bool {
	// get min and max bounds of rectangle
	rectMinX, rectMaxX := MinMax(rect.start.x, rect.end.x)
	rectMinY, rectMaxY := MinMax(rect.start.y, rect.end.y)

	var currEdge TilePair
	var edgeMinX, edgeMinY, edgeMaxX, edgeMaxY, i int
	var currEdgeRange Range

	// test all horizontal loop edges for collisions with rectangle
	for yCoord, currEdges := range horizontalEdges {
		for i = range len(currEdges) {
			currEdge = currEdges[i]

			edgeMinX, edgeMaxX = MinMax(currEdge.start.x, currEdge.end.x)
			currEdgeRange = Range{edgeMinX, edgeMaxX}

			rectHasLeftOrRightEdgeInRange := InRange(rectMinX+1, currEdgeRange) || InRange(rectMaxX-1, currEdgeRange)
			edgeInRectYRange := InRange(yCoord, Range{rectMinY + 1, rectMaxY - 1})

			if rectHasLeftOrRightEdgeInRange && edgeInRectYRange {
				return false
			}
		}
	}

	// test vertical edges of rectangle
	for xCoord, currEdges := range verticalEdges {
		for i = range len(currEdges) {
			currEdge = currEdges[i]

			edgeMinY, edgeMaxY = MinMax(currEdge.start.y, currEdge.end.y)
			currEdgeRange = Range{edgeMinY, edgeMaxY}

			rectHasTopOrBottomEdgeInRange := InRange(rectMinY+1, currEdgeRange) || InRange(rectMaxY-1, currEdgeRange)
			edgeInRectXRange := InRange(xCoord, Range{rectMinX + 1, rectMaxX - 1})

			if rectHasTopOrBottomEdgeInRange && edgeInRectXRange {
				return false
			}
		}
	}

	// check if arbitrary point of rect is inside loop
	avgRectX, avgRectY := (rectMinX+rectMaxX)/2, (rectMinY+rectMaxY)/2

	if !TileInsideLoop(Vec2{avgRectX, avgRectY}) {
		return false
	}

	return true
}

// Determines whether the tile with the given coordinates is within (or on the edge of) the loop of red/green tiles
func TileInsideLoop(tile Vec2) bool {
	// check cacche to see if result is already computed
	cachedResult, exists := insideLoopCache[tile]
	if exists {
		return cachedResult
	}

	// check if the tile itself is on an edge
	if TileOnEdge(tile, true) || TileOnEdge(tile, false) {
		insideLoopCache[tile] = true
		return true
	}

	// count the number of vertical edges encountered between tile and minimum edge of relevant tile range
	edgesEncountered := 0
	for i := tile.x - 1; i >= minX; i-- {
		if TileOnEdge(Vec2{i, tile.y}, false) {
			edgesEncountered++
		}
	}

	// the tile is inside the loop if the number of edges between itself and the edge of the range of relevant tiles is odd
	result := edgesEncountered%2 != 0

	insideLoopCache[tile] = result
	return result
}

// Determines whether the tile with the given coordinates is on one of the edges
func TileOnEdge(tile Vec2, horizontal bool) bool {
	var edgeMap map[int][]TilePair
	var coordKey, testCoord int

	if horizontal {
		edgeMap = horizontalEdges
		coordKey = tile.y
		testCoord = tile.x
	} else {
		edgeMap = verticalEdges
		coordKey = tile.x
		testCoord = tile.y
	}

	currEdges, exists := edgeMap[coordKey]

	if !exists {
		return false
	}

	var currEdge TilePair
	var upperBound, lowerBound int
	for i := range len(currEdges) {
		currEdge = currEdges[i]

		if horizontal {
			// horizontal edge -> test tile's x coordinate
			lowerBound, upperBound = MinMax(currEdge.start.x, currEdge.end.x)
		} else {
			// vertical edge -> test tile's y coordinate
			lowerBound, upperBound = MinMax(currEdge.start.y, currEdge.end.y)
		}

		if testCoord >= lowerBound && testCoord <= upperBound {
			return true
		}
	}

	return false
}

// Returns the minimum and maximum of the two given numbers in the format of (min, max).
func MinMax(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num2, num1
	} else {
		return num1, num2
	}
}
