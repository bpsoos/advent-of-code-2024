package day8

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	_, err = os.ReadFile("inputs/day8_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 8 test 1:")
	fmt.Println()
	fmt.Println("day 8 solution 1:")
	fmt.Println()
	fmt.Println("day 8 test 2:")
	fmt.Println()
	fmt.Println("day 8 solution 2:")
	fmt.Println()
}
