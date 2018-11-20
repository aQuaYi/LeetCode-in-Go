package problem0911

import "sort"

// TopVotedCandidate object will be instantiated and called as such:
// obj := Constructor(persons, times);
// param_1 := obj.Q(t);
type TopVotedCandidate struct {
	times   []int
	leaders []int
}

// Constructor returns TopVotedCandidate
func Constructor(persons []int, times []int) TopVotedCandidate {
	size := len(persons)

	votes := make([]int, size)
	leaders := make([]int, size)

	leader := persons[0]
	for i := 0; i < size; i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}
	return TopVotedCandidate{times, leaders}
}

// Q is quest
func (tvc *TopVotedCandidate) Q(t int) int {
	i := sort.SearchInts(tvc.times, t)
	if i == len(tvc.times) || t != tvc.times[i] {
		i--
	}
	return tvc.leaders[i]
}
