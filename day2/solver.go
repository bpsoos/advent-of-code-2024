package day2

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day2_test.txt")
	if err != nil {
		panic(err)
	}
	input, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 2 test 1:")
	fmt.Println(CountSafeReports(ParseReports(bytes.NewReader(testInput))))

	fmt.Println("day 2 solution 1:")
	fmt.Println(CountSafeReports(ParseReports(bytes.NewReader(input))))

	fmt.Println("day 2 test 2:")
	fmt.Println(CountSafeReportsDampened(ParseReports(bytes.NewReader(testInput))))

	fmt.Println("day 2 solution 2:")
	fmt.Println(CountSafeReportsDampened(ParseReports(bytes.NewReader(input))))
}
