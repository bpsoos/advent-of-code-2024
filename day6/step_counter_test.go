package day6_test

import (
	"advent2024/day6"
	"strings"
	"testing"
)

func TestStepCounterGivenTest(T *testing.T) {
	rawLabMap := `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != 41 {
		T.Fatalf("result (%d) does not match expected (%d)", result, 41)
	}
}

func TestStepCounterSimpleRightEnd(T *testing.T) {
	rawLabMap := `
....#.
......
.#..^.
......
`
	want := 3
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestStepCounterSimpleTopEnd(T *testing.T) {
	rawLabMap := `
.#....
.....#
.^#...
....#.
`
	want := 8
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestStepCounterSimpleLeftEnd(T *testing.T) {
	rawLabMap := `
.#....
.....#
.^....
....#.
`
	want := 9
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestStepCounterSimpleBottomEnd(T *testing.T) {
	rawLabMap := `
.#....
.....#
.^....
......
`
	want := 7
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestStepCounterTripleTurn(T *testing.T) {
	rawLabMap := `
.#....
.....#
.^.#.#
....#.
`
	want := 7
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestStepCounterTripleTurnLeft(T *testing.T) {
	rawLabMap := `
......
...#..
.#...#
...^..
.#....
..#.#.
......
`
	want := 11
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).Count()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithGivenTest(T *testing.T) {
	rawLabMap := `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	want := 6
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithCrossingOwnPath(T *testing.T) {
	rawLabMap := `
.#....
.....#
......
......
....#.
.^....
......
`
	want := 1
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithCrossingOwnPathColumn(T *testing.T) {
	rawLabMap := `
.#....
.....#
......
......
.^....
......
....#.
`
	want := 1
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithTurningAtOwnPathTraceBlocked(T *testing.T) {
	rawLabMap := `
.#..#.
.....#
......
......
.^....
......
....#.
`
	want := 2
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithCrossingLoopRow(T *testing.T) {
	rawLabMap := `
......
.#....
#....#
....#.
..^...
......
......
`
	want := 1
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithCrossingLoopColumn(T *testing.T) {
	rawLabMap := `
...#..
.#..#.
......
......
..#...
.^.#..
......
`
	want := 1
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}

func TestLoopCountWithLoopAtStartPosition(T *testing.T) {
	rawLabMap := `
......
.##...
...#..
......
......
.^....
..#...
`
	want := 1
	result := day6.NewStepCounter(day6.ParseMap(strings.NewReader(rawLabMap))).CountLoop()
	if result != want {
		T.Fatalf("result (%d) does not match expected (%d)", result, want)
	}
}
