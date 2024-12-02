package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Levels []int

type Report []Levels

func ParseReports(rawReport io.Reader) Report {
	report := make(Report, 0)
	scanner := bufio.NewScanner(rawReport)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		splitLine := strings.Split(scanner.Text(), " ")
		levels := make(Levels, len(splitLine))
		for i, rawLevel := range splitLine {
			level, err := strconv.Atoi(rawLevel)
			if err != nil {
				panic(err)
			}
			levels[i] = level
		}
		report = append(report, levels)
	}

	return report
}
