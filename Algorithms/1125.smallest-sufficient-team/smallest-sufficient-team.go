package problem1125

import "math"

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	index := make(map[string]uint)
	for i, skill := range reqSkills {
		index[skill] = uint(i)
	}
	// encode person's skills to integer
	personSkills := make([]int, len(people))
	// collect all candidates of a skill
	cands := make([][]int, len(reqSkills))
	for i, p := range people {
		skills := 0
		for _, skill := range p {
			if _, ok := index[skill]; ok {
				skills |= 1 << index[skill]
				cands[index[skill]] = append(cands[index[skill]], i)
			}
		}
		personSkills[i] = skills
	}

	smallestTeam := map[int][]int{0: {}}

	var dfs func(int) []int
	dfs = func(skills int) []int {
		if team, ok := smallestTeam[skills]; ok {
			return team
		}
		var minTeam []int
		minTeamSize := 61 // people.length<=60
		skill := int(math.Log2(float64(skills & -skills)))
		for _, c := range cands[skill] {
			needSkills := skills & (^personSkills[c])
			team := dfs(needSkills)
			if minTeamSize > len(team)+1 {
				minTeamSize = len(team) + 1
				minTeam = append(team, c)
			}
		}
		smallestTeam[skills] = minTeam
		return minTeam
	}

	return dfs(1<<uint(len(reqSkills)) - 1)
}
