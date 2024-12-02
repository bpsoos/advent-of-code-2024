package day1

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func ParseLocations(rawLocations io.Reader) ([]int, []int) {
	scanner := bufio.NewScanner(rawLocations)
	group1 := make([]int, 0)
	group2 := make([]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		splitLine := strings.Split(scanner.Text(), "  ")
		n1, err := strconv.Atoi(strings.TrimSpace(splitLine[0]))
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(strings.TrimSpace(splitLine[1]))
		if err != nil {
			panic(err)
		}
		group1 = append(group1, n1)
		group2 = append(group2, n2)
	}

    sort.Ints(group1)
    sort.Ints(group2)

	return group1, group2
}
