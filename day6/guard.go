package day6

type Guard struct {
	Position
	Direction Direction
}

func (g *Guard) TurnRightDir() Direction {
	switch g.Direction {
	case DirectionUp:
		return DirectionRight
	case DirectionRight:
		return DirectionDown
	case DirectionDown:
		return DirectionLeft
	case DirectionLeft:
		return DirectionUp
	default:
		panic("invalid direction")
	}
}

func (g *Guard) TurnRight() {
	switch g.Direction {
	case DirectionUp:
		g.Direction = DirectionRight
	case DirectionRight:
		g.Direction = DirectionDown
	case DirectionDown:
		g.Direction = DirectionLeft
	case DirectionLeft:
		g.Direction = DirectionUp
	default:
		panic("invalid direction")
	}
}

func (g Guard) DeepCopy() *Guard {
	return &Guard{
		Position:  g.Position.DeepCopy(),
		Direction: g.Direction,
	}
}
