package day5

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseUpdateInfo(rawInfo io.Reader) (Rules, []Update) {
	rules := make(Rules)
	updates := make([]Update, 0)
	scanner := bufio.NewScanner(rawInfo)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if strings.Contains(scanner.Text(), "|") {
			splitLine := strings.Split(scanner.Text(), "|")
			x, err := strconv.Atoi(splitLine[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(splitLine[1])
			if err != nil {
				panic(err)
			}
			_, ok := rules[x]
			if !ok {
				rules[x] = []int{y}
			} else {
				rules[x] = append(rules[x], y)
			}
			continue
		}
		splitLine := strings.Split(scanner.Text(), ",")
		update := make(Update, len(splitLine))
		for i := range splitLine {
			pageNumber, err := strconv.Atoi(splitLine[i])
			if err != nil {
				panic(err)
			}
			update[i] = pageNumber
		}
		updates = append(updates, update)
	}

	return rules, updates
}
