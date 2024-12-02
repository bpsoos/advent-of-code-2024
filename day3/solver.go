package day3

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 3 test 1:")
	fmt.Println()
	fmt.Println("day 3 solution 1:")
	fmt.Println()
	fmt.Println("day 3 test 2:")
	fmt.Println()
	fmt.Println("day 3 solution 2:")
	fmt.Println()
}
