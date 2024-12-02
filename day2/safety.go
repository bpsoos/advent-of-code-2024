package day2

import (
	"math"
)

func (l Levels) IsSafe() bool {
	asc := l[0] < l[1]

	for i := 1; i < len(l); i++ {
		if l[i-1] == l[i] {
			return false
		}
		ascCurrent := l[i-1] < l[i]
		if asc != ascCurrent {
			return false
		}
		if math.Abs(float64(l[i]-l[i-1])) > 3 {
			return false
		}
	}

	return true
}

func CountSafeReports(report Report) int {
	count := 0
	for _, levels := range report {
		if levels.IsSafe() {
			count++
		}
	}

	return count
}
