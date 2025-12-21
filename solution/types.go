package solution

type SolutionPair interface {
	Init(string)
	PartOne() string
	PartTwo() string
}

type Range struct {
	low  int
	high int
}
