package day4

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	input, err := os.ReadFile("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	testInput, err := os.ReadFile("inputs/day4_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 4 test 1:")
	fmt.Println(SearchXmas(bytes.NewReader(testInput)))
	fmt.Println("day 4 solution 1:")
	fmt.Println(SearchXmas(bytes.NewReader(input)))
	fmt.Println("day 4 test 2:")
	fmt.Println(SearchMasX(bytes.NewReader(testInput)))
	fmt.Println("day 4 solution 2:")
	fmt.Println(SearchMasX(bytes.NewReader(input)))
}
