package day1

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day1_test.txt")
	fmt.Println("day 1 test solution 1:")
	l1, l2 := ParseLocations(bytes.NewReader(testInput))
	fmt.Printf("%v %v\n", l1, l2)
	fmt.Println(getSolution1(l1, l2))

	input, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 1 solution 1:")
	fmt.Println(getSolution1(ParseLocations(bytes.NewReader(input))))

	fmt.Println("day 1 test solution 2:")
	fmt.Println(getSolution2(ParseLocations(bytes.NewReader(testInput))))
	fmt.Println("day 1 solution 2:")
	fmt.Println(getSolution2(ParseLocations(bytes.NewReader(input))))
}

func getSolution1(group1 []int, group2 []int) int {
	count := 0
	for i := range len(group1) {
		sum := group1[i] - group2[i]
		if sum < 0 {
			sum = -1 * sum
		}
		count += sum
	}

	return count
}

func getSolution2(group1 []int, group2 []int) int {
	score := 0
	for i := range len(group1) {
		count := 0
		for j := range len(group1) {
			if group2[j] == group1[i] {
				count++
			}
		}
		score += count * group1[i]
	}

	return score
}
