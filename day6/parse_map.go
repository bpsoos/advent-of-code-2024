package day6

import (
	"bufio"
	"io"
)

func ParseMap(rawMap io.Reader) *LabMap {
	scanner := bufio.NewScanner(rawMap)
	labMap := new(LabMap)

	i := 0
	width := 0
	positions := make([][]Position, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		width = len(scanner.Text())
		posRow := make([]Position, 0)
		for j, char := range scanner.Text() {
			pos := NewPosition(j, i)
			if char == '#' {
				pos.IsObstruction = true
			}
			if char == '^' {
				dir := DirectionUp
				pos.VisitDirections = []Direction{DirectionUp}
				labMap.StepCount = 1
				labMap.startPosition = *pos
				labMap.guard = &Guard{
					Position:  *pos,
					Direction: dir,
				}
			}
			posRow = append(posRow, *pos)
		}
		positions = append(positions, posRow)
		i++
	}

	labMap.positions = positions
	labMap.width = width
	labMap.height = i

	return labMap
}
