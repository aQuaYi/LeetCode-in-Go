package problem1125

import "math"

var workers []int
var skillCands [][]int
var dp map[int][]int

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	index := make(map[string]int)
	for i, skill := range reqSkills {
		index[skill] = i
	}
	workers = make([]int, len(people))
	skillCands = make([][]int, len(reqSkills))
	for i, p := range people {
		for _, skill := range p {
			if _, ok := index[skill]; ok {
				workers[i] |= (1 << uint(index[skill]))
				skillCands[index[skill]] = append(skillCands[index[skill]], i)
			}
		}
	}
	dp = map[int][]int{0: {}}
	return dfs(1<<uint(len(reqSkills)) - 1)
}

func dfs(skillsWanted int) []int {
	if v, ok := dp[skillsWanted]; ok {
		return v
	}
	smallestTeamSize := 61
	res := []int{}
	targetSkill := int(math.Log2(float64(skillsWanted & -skillsWanted)))
	for _, p := range skillCands[targetSkill] {
		candSkillWanted := skillsWanted &^ workers[p]
		pre := dfs(candSkillWanted)
		pre = append(pre, p)
		if len(pre) < smallestTeamSize {
			smallestTeamSize = len(pre)
			res = pre
		}
	}
	dp[skillsWanted] = res
	return res
}
