package day6

import (
	"bufio"
	"io"
)

func ParseMap(rawMap io.Reader) *LabMap {
	scanner := bufio.NewScanner(rawMap)
	labMap := new(LabMap)

	i := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		row := make([]Position, 0)
		for j, char := range scanner.Text() {
			pos := Position{
				IsObstruction: false,
				X:             j,
				Y:             i,
			}
			if char == '#' {
				pos.IsObstruction = true
			}
			if char == '^' {
				labMap.guard = &Guard{
					Position:  pos,
					Direction: UpDirection,
				}
			}
			row = append(row, pos)
		}
		labMap.positions = append(labMap.positions, row)
	}

	return labMap
}
