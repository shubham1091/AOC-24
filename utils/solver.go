package utils

// DaySolver defines the interface for daily solutions
type DaySolver interface {
	Solve() (interface{}, error)
	PartOne(input []string) (interface{}, error)
	PartTwo(input []string) (interface{}, error)
}
