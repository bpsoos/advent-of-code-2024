package day7

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.ReadFile("inputs/day7_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 7 test 1:")
	fmt.Println()
	fmt.Println("day 7 solution 1:")
	fmt.Println()
	fmt.Println("day 7 test 2:")
	fmt.Println()
	fmt.Println("day 7 solution 2:")
	fmt.Println()
}
