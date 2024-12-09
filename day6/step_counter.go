package day6

import (
	"slices"
)

type StepCounter struct {
	incrementer *Incrementer
	labMap      *LabMap
}

func NewStepCounter(labMap *LabMap) *StepCounter {
	return &StepCounter{
		incrementer: &Incrementer{
			Width:  labMap.width,
			Height: labMap.height,
		},
		labMap: labMap,
	}
}

func (sc *StepCounter) Count() int {
	for sc.labMap.Update() {
	}

	return sc.labMap.StepCount
}

func (sc *StepCounter) CountLoop() int {
	obstructions := make([]Position, 0)

	for sc.incrementer.Increment() {
		obstructions = sc.simulateObs(obstructions)
	}

	return len(obstructions)
}

func (sc *StepCounter) simulateObs(obs []Position) []Position {
	altMap := sc.generateAlternativeMap()
	if altMap == nil {
		return obs
	}

	obsCandidate := altMap.FindPosition(sc.incrementer.X, sc.incrementer.Y)
	if sc.isValidCandidate(obsCandidate, obs) {
		return obs
	}

	for altMap.Update() {
	}
	if altMap.GuardIsAtEdge() {
		return obs
	}

	obs = append(obs, *obsCandidate)

	return obs
}

func (sc *StepCounter) isValidCandidate(candidatePos *Position, obs []Position) bool {
	return candidatePos == nil ||
		slices.ContainsFunc(
			obs,
			func(pos Position) bool { return pos.Eq(*candidatePos) },
		) ||
		candidatePos.Eq(sc.labMap.startPosition)
}

func (sc *StepCounter) generateAlternativeMap() *LabMap {
	altMap := sc.labMap.DeepCopy()
	pos := sc.labMap.FindPosition(sc.incrementer.X, sc.incrementer.Y)
	if pos == nil || pos.IsObstruction {
		return nil
	}
	altMap.FindPosition(sc.incrementer.X, sc.incrementer.Y).IsObstruction = true

	return altMap
}

type Incrementer struct {
	X, Y, Width, Height int
}

func (i *Incrementer) Increment() bool {
	if i.X == i.Width && i.Y != i.Height {
		i.X = 0
		i.Y++
		return true
	}

	if i.X == i.Width && i.Y == i.Height {
		i.X = 0
		i.Y = 0
		return false
	}

	i.X++
	return true
}
