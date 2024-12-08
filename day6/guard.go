package day6

type Guard struct {
	Position
	Direction Direction
}

func (g *Guard) TurnRight() {
	if g.Direction == LeftDirection {
		g.Direction = UpDirection
		return
	}
	g.Direction++
}

func (g *Guard) Step() {
	switch g.Direction {
	case UpDirection:
		g.Y--
	case DownDirection:
		g.Y++
	case LeftDirection:
		g.X--
	case RightDirection:
		g.X++
	}
}
