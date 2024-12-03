package day3

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day3_test.txt")
	if err != nil {
		panic(err)
	}
	input, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 3 test 1:")
	fmt.Println(calculate(ParseMemory(bytes.NewReader(testInput))))
	fmt.Println("day 3 solution 1:")
	fmt.Println(calculate(ParseMemory(bytes.NewReader(input))))

	test2Input, err := os.ReadFile("inputs/day3_test2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 3 test 2:")
	fmt.Println(calculate(ParseMemoryWithDosAndDonts(bytes.NewReader(test2Input))))
	fmt.Println("day 3 solution 2:")
	fmt.Println(calculate(ParseMemoryWithDosAndDonts(bytes.NewReader(input))))
}

func calculate(muls []Mul) int {
	result := 0
	for _, mul := range muls {
		result += mul[0] * mul[1]
	}

	return result
}
