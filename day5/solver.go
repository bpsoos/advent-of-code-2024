package day5

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.ReadFile("inputs/day5_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 5 test 1:")
	fmt.Println()
	fmt.Println("day 5 solution 1:")
	fmt.Println()
	fmt.Println("day 5 test 2:")
	fmt.Println()
	fmt.Println("day 5 solution 2:")
	fmt.Println()
}
