package main

import (
	"advent2024/day1"
	"advent2024/day2"
	"advent2024/day3"
	"advent2024/day4"
	"advent2024/day5"
	"advent2024/day6"
	"advent2024/day7"
	"fmt"
	"os"
)

type Solver interface {
	Solve()
}

func main() {
	day := os.Args[1]
	solvers := map[string]Solver{
		"day1": day1.Solver{},
		"day2": day2.Solver{},
		"day3": day3.Solver{},
		"day4": day4.Solver{},
		"day5": day5.Solver{},
		"day6": day6.Solver{},
		"day7": day7.Solver{},
	}
	solver, ok := solvers[day]
	if !ok {
		fmt.Println("solver not found for")
		fmt.Println(day)
		os.Exit(1)
	}
	solver.Solve()
}
