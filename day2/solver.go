package day2

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 2 solution 1:")
	fmt.Println()
	fmt.Println("day 2 solution 2:")
	fmt.Println()
}

