package day7

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	input, err := os.ReadFile("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	testInput, err := os.ReadFile("inputs/day7_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 7 test 1:")
	fmt.Println(calculateOne(ParseEquations(bytes.NewReader(testInput))))
	fmt.Println("day 7 solution 1:")
	fmt.Println(calculateOne(ParseEquations(bytes.NewReader(input))))
	fmt.Println("day 7 test 2:")
	fmt.Println(calculateTwo(ParseEquations(bytes.NewReader(testInput))))
	fmt.Println("day 7 solution 2:")
	fmt.Println(calculateTwo(ParseEquations(bytes.NewReader(input))))
}

func calculateOne(eqs []CalibrationEquation) uint64 {
	sum := uint64(0)
	evaluator := EquationEvaluator{}

	for _, eq := range eqs {
		if evaluator.Evaluate(&eq) {
			sum += eq.TestValue
		}
	}

	return sum
}

func calculateTwo(eqs []CalibrationEquation) uint64 {
	sum := uint64(0)
	evaluator := EquationEvaluator{}

	for _, eq := range eqs {
		if evaluator.EvaluateWithThreeOperators(&eq) {
			sum += eq.TestValue
		}
	}

	return sum
}
