package day6

type Position struct {
	IsObstruction bool
	X             int
	Y             int
}

type Direction int

const (
	UpDirection    Direction = iota
	RightDirection           = iota
	DownDirection            = iota
	LeftDirection            = iota
)

type LabMap struct {
	positions [][]Position
	guard     *Guard
}

func (lm *LabMap) Guard() *Guard {
	return lm.guard
}

func (lm *LabMap) FacingWall(guard *Guard) bool {
	return false
}

func (lm *LabMap) Update() {
	if lm.FacingWall(lm.guard) {
		lm.guard.TurnRight()
	}
	lm.guard.Step()
}

func (lm *LabMap) IsAtEdge(position *Position) bool {
	if position.X == 0 || position.Y == 0 {
		return true
	}

	if position.X == len(lm.positions[0])-1 || position.Y == len(lm.positions)-1 {
		return true
	}

	return false
}
