package day1

import (
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	_, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 1 solution 1:")
	fmt.Println("day 1 solution 2:")
}
