package day6

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	input, err := os.ReadFile("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	testInput, err := os.ReadFile("inputs/day6_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 6 test 1:")
	fmt.Println(NewStepCounter(ParseMap(bytes.NewReader(testInput))).Count())
	fmt.Println("day 6 solution 1:")
	fmt.Println(NewStepCounter(ParseMap(bytes.NewReader(input))).Count())
	fmt.Println("day 6 test 2:")
	fmt.Println(NewStepCounter(ParseMap(bytes.NewReader(testInput))).CountLoop())
	fmt.Println("day 6 solution 2:")
	fmt.Println(NewStepCounter(ParseMap(bytes.NewReader(input))).CountLoop())
}
