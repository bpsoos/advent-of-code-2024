package day6

type Direction int

const DirectionUnset Direction = -1
const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type Position struct {
	IsObstruction   bool
	X               int
	Y               int
	VisitDirections []Direction
}

func (p Position) DeepCopy() Position {
	visitDirections := make([]Direction, len(p.VisitDirections))
	copy(visitDirections, p.VisitDirections)

	return Position{
		IsObstruction:   p.IsObstruction,
		X:               p.X,
		Y:               p.Y,
		VisitDirections: visitDirections,
	}
}

func NewPosition(x, y int) *Position {
	return &Position{
		X:               x,
		Y:               y,
		VisitDirections: make([]Direction, 0),
	}
}

func (p Position) Eq(pos Position) bool {
	return p.X == pos.X &&
		p.Y == pos.Y
}
