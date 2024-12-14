package day6

import (
	"slices"
)

type LabMap struct {
	guard         *Guard
	StepCount     int
	positions     [][]Position
	startPosition Position
	width         int
	height        int
}

func (lm *LabMap) Update() bool {
	if lm.GuardIsAtEdge() {
		return false
	}

	nextPos, nextDir := lm.findNextGuardStep()
	if lm.isLoopingPos(*nextPos, nextDir) {
		return false
	}

	lm.moveGuard(nextPos, nextDir)

	return true
}

func (lm *LabMap) moveGuard(pos *Position, dir Direction) {
	lm.guard.Position = *pos
	lm.guard.Direction = dir
	if len(pos.VisitDirections) == 0 {
		lm.StepCount++
	}
	pos.VisitDirections = append(pos.VisitDirections, dir)
}

func (lm *LabMap) isEdge(pos Position) bool {
	return pos.X == 0 ||
		pos.X == lm.width-1 ||
		pos.Y == 0 ||
		pos.Y == lm.height-1
}

func (lm *LabMap) GuardIsAtEdge() bool {
	return lm.isEdge(lm.guard.Position)
}

func (lm *LabMap) isLoopingPos(pos Position, dir Direction) bool {
	return slices.Contains(pos.VisitDirections, dir)
}

func (lm *LabMap) ForwardStep() *Position {
	x := lm.guard.X
	y := lm.guard.Y

	switch lm.guard.Direction {
	case DirectionUp:
		y--
	case DirectionRight:
		x++
	case DirectionDown:
		y++
	case DirectionLeft:
		x--
	default:
		panic("invalid direction")
	}

	return lm.FindPosition(x, y)
}

func (lm *LabMap) LeftStep() *Position {
	x := lm.guard.X
	y := lm.guard.Y

	switch lm.guard.Direction {
	case DirectionUp:
		x--
	case DirectionRight:
		y++
	case DirectionDown:
		x++
	case DirectionLeft:
		y--
	default:
		panic("invalid direction")
	}

	return lm.FindPosition(x, y)
}

func (lm *LabMap) RightStep() *Position {
	x := lm.guard.X
	y := lm.guard.Y

	switch lm.guard.Direction {
	case DirectionUp:
		x++
	case DirectionRight:
		y--
	case DirectionDown:
		x--
	case DirectionLeft:
		y++
	default:
		panic("invalid direction")
	}

	return lm.FindPosition(x, y)
}

func (lm *LabMap) BackStep() *Position {
	x := lm.guard.X
	y := lm.guard.Y

	switch lm.guard.Direction {
	case DirectionUp:
		y++
	case DirectionRight:
		x--
	case DirectionDown:
		y--
	case DirectionLeft:
		x--
	default:
		panic("invalid direction")
	}

	return lm.FindPosition(x, y)
}

func (lm *LabMap) findNextGuardStep() (*Position, Direction) {
	nextPosCandidate := lm.ForwardStep()
	if !nextPosCandidate.IsObstruction {
		return nextPosCandidate, lm.guard.Direction
	}

	return lm.FindPosition(lm.guard.Position.X, lm.guard.Position.Y), lm.guard.TurnRightDir()
}

func (lm *LabMap) FindPosition(x, y int) *Position {
	if x >= lm.width || y >= lm.height || x < 0 || y < 0 {
		return nil
	}

	return &lm.positions[y][x]
}

func (lm LabMap) DeepCopy() *LabMap {
	positions := make([][]Position, len(lm.positions))
	for i := range positions {
		positions[i] = make([]Position, len(lm.positions[0]))
		for j := range lm.positions[0] {
			positions[i][j] = lm.positions[i][j].DeepCopy()
		}
	}

	return &LabMap{
		guard:     lm.guard.DeepCopy(),
		StepCount: lm.StepCount,
		positions: positions,
		width:     lm.width,
		height:    lm.height,
	}
}
