package day6

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.ReadFile("inputs/day6_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 6 test 1:")
	fmt.Println()
	fmt.Println("day 6 solution 1:")
	fmt.Println()
	fmt.Println("day 6 test 2:")
	fmt.Println()
	fmt.Println("day 6 solution 2:")
	fmt.Println()
}
