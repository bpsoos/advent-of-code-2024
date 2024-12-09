package day7

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	testInput, err = os.ReadFile("inputs/day7_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 7 test 1:")
	ParseEquations(bytes.NewReader(testInput))
	fmt.Println()
	fmt.Println("day 7 solution 1:")
	fmt.Println()
	fmt.Println("day 7 test 2:")
	fmt.Println()
	fmt.Println("day 7 solution 2:")
	fmt.Println()
}


