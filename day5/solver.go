package day5

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	input, err := os.ReadFile("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	testInput, err := os.ReadFile("inputs/day5_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 5 test 1:")
	rules, updates := ParseUpdateInfo(bytes.NewReader(testInput))
	fmt.Println(calculateUpdateInfo(filterIncorrectUpdates(updates, rules)))
	fmt.Println("day 5 solution 1:")
	rules, updates = ParseUpdateInfo(bytes.NewReader(input))
	fmt.Println(calculateUpdateInfo(filterIncorrectUpdates(updates, rules)))
	fmt.Println("day 5 test 2:")
	rules, updates = ParseUpdateInfo(bytes.NewReader(testInput))
	fmt.Println(calculateUpdateInfo(sortIncorrectUpdates(filterCorrectUpdates(updates, rules), rules)))
	fmt.Println("day 5 solution 2:")
	rules, updates = ParseUpdateInfo(bytes.NewReader(input))
	fmt.Println(calculateUpdateInfo(sortIncorrectUpdates(filterCorrectUpdates(updates, rules), rules)))
}

func calculateUpdateInfo(verifiedUpdates []Update) int {
	count := 0
	for _, update := range verifiedUpdates {
		count += update[int(len(update)/2)]
	}
	return count
}
