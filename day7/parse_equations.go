package day7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CalibrationEquation struct {
	TestValue uint64
	Equation  []int
}

func ParseEquations(rawEquations io.Reader) []CalibrationEquation {
	scanner := bufio.NewScanner(rawEquations)
	equations := make([]CalibrationEquation, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		splitLine := strings.Split(scanner.Text(), ":")
		testValue, err := strconv.ParseUint(splitLine[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rawEquation := strings.Split(strings.TrimSpace(splitLine[1]), " ")
		equation := make([]int, len(rawEquation))
		for i := range rawEquation {
			num, err := strconv.Atoi(rawEquation[i])
			if err != nil {
				panic(err)
			}
			equation[i] = num
		}

		equations = append(equations,
			CalibrationEquation{
				TestValue: testValue,
				Equation:  equation,
			})
	}

	fmt.Printf("equations: %v\n", equations)
	return equations
}
