package main

import (
	"advent2024/day1"
	"advent2024/day2"
	"fmt"
	"os"
)

type Solver interface {
	Solve()
}

func main() {
	day := os.Args[1]
	solvers := map[string]Solver{
		"day1":  day1.Solver{},
		"day2":  day2.Solver{},
	}
	solver, ok := solvers[day]
	if !ok {
		fmt.Println("solver not found for")
		fmt.Println(day)
		os.Exit(1)
	}
	solver.Solve()
}
