package day7

import (
	"math"
	"strconv"
)

type EquationEvaluator struct{}

func (od *EquationEvaluator) Evaluate(ce *CalibrationEquation) bool {
	return evalEquation(2, ce, updateVal)
}

func (od *EquationEvaluator) EvaluateWithThreeOperators(ce *CalibrationEquation) bool {
	return evalEquation(3, ce, updateValWithThreeOperators)
}

func evalEquation(
	base int,
	ce *CalibrationEquation,
	updateVal func(operator rune, val uint64, nextVal uint64) uint64,
) bool {
	maxPermutations := int(math.Pow(float64(base), float64(len(ce.Equation)-1)))

	for i := 0; i < maxPermutations; i++ {
		val := ce.Equation[0]
		s := strconv.FormatUint(uint64(i), base)
		for range len(ce.Equation) - 1 - len(s) {
			s = "0" + s
		}

		for j, b := range s {
			val = updateVal(b, val, ce.Equation[j+1])
		}

		if val == ce.TestValue {
			return true
		}
	}

	return false
}

func updateVal(operator rune, val uint64, nextVal uint64) uint64 {
	if operator == '0' {
		return val * nextVal
	}

	return val + nextVal
}

func updateValWithThreeOperators(operator rune, val uint64, nextVal uint64) uint64 {
	if operator == '0' {
		return val * nextVal
	}
	if operator == '1' {
		return val + nextVal
	}
	if operator == '2' {
		val, err := strconv.ParseUint(strconv.FormatUint(val, 10)+strconv.FormatUint(nextVal, 10), 0, 64)

		if err != nil {
			panic(err)
		}

		return val
	}
	panic("unknown operator")

}
