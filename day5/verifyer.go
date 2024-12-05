package day5

import (
	"slices"
	"sort"
)

type Update []int
type Rules map[int][]int

type SortableUpdate struct {
	Update Update
	Rules  Rules
}

type Verifier struct {
	Rules Rules
}

func (su SortableUpdate) Len() int      { return len(su.Update) }
func (su SortableUpdate) Swap(i, j int) { su.Update[i], su.Update[j] = su.Update[j], su.Update[i] }
func (su SortableUpdate) Less(i, j int) bool {
	ruleSet, ok := su.Rules[su.Update[i]]
	if !ok {
		return false
	}

	return slices.Contains(ruleSet, su.Update[j])
}

func (v Verifier) Verify(update Update) bool {
	for i := range update {
		ruleSet, ok := v.Rules[update[i]]
		if !ok {
			continue
		}
		for _, rule := range ruleSet {
			if slices.Contains[[]int](update[:i], rule) {
				return false
			}
		}
	}
	return true
}

func filterIncorrectUpdates(updates []Update, rules Rules) []Update {
	verifiedUpdates := make([]Update, 0)
	verifier := Verifier{Rules: rules}
	for _, update := range updates {
		if verifier.Verify(update) {
			verifiedUpdates = append(verifiedUpdates, update)
		}
	}

	return verifiedUpdates
}

func filterCorrectUpdates(updates []Update, rules Rules) []Update {
	verifiedUpdates := make([]Update, 0)
	verifier := Verifier{Rules: rules}
	for _, update := range updates {
		if !verifier.Verify(update) {
			verifiedUpdates = append(verifiedUpdates, update)
		}
	}

	return verifiedUpdates
}

func sortIncorrectUpdates(updates []Update, rules Rules) []Update {
	sorted := make([]Update, len(updates))
	for i := range updates {
		sortable := SortableUpdate{Update: updates[i], Rules: rules}
		sort.Sort(sortable)
		sorted[i] = sortable.Update
	}

	return sorted
}
