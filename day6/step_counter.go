package day6

type StepCounter struct {
	steps []Position
}

func NewStepCounter() *StepCounter {
	return &StepCounter{steps: make([]Position, 0)}
}

func (sc *StepCounter) Count(labMap *LabMap) int {
	sc.observePatrol(labMap)
	return len(sc.steps)
}

func (sc *StepCounter) observePatrol(labMap *LabMap) {
	for {
		guard := labMap.Guard()
		sc.steps = append(sc.steps, guard.Position)
		if labMap.IsAtEdge(&guard.Position) {
			return
		}
		labMap.Update()
	}
}
