package day4

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.ReadFile("inputs/day4_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 4 test 1:")
	fmt.Println()
	fmt.Println("day 4 solution 1:")
	fmt.Println()
	fmt.Println("day 4 test 2:")
	fmt.Println()
	fmt.Println("day 4 solution 2:")
	fmt.Println()
}
