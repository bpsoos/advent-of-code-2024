package day7

import "fmt"

type EquationEvaluator struct{}

func (od *EquationEvaluator) Evaluate(ce *CalibrationEquation) bool {
	operatorFlips := make([]bool, len(ce.Equation)-1)
	return evalEquation(operatorFlips, ce)
}

func evalEquation(opFlips []bool, ce *CalibrationEquation) bool {
	for i := 0; i < len(opFlips)*len(opFlips); i++ {
		val := uint64(ce.Equation[0])
		s := fmt.Sprintf("%b", i)
		for range len(opFlips) - len(s) {
			s = "0" + s
		}

		for j, b := range s {
			val = updateVal(b == '0', val, ce.Equation[j+1])
		}

		if val == ce.TestValue {
			return true
		}
	}

	return false
}

func updateVal(flip bool, val uint64, nextVal int) uint64 {
	if flip {
		return val * uint64(nextVal)
	}

	return val + uint64(nextVal)
}

type Permutator struct {
	permutable []bool
	i          int
	j          int
}

func NewPermutator(size int) *Permutator {
	return &Permutator{
		permutable: make([]bool, size),
		i:          size * size,
		j:          size,
	}
}

func (p *Permutator) Next() bool {
	return p.i < 0
}

func (p Permutator) Get() []bool {
	return p.permutable
}
