package problem1125

import "math"

func smallestSufficientTeam(reqs []string, peoples [][]string) []int {
	n := len(reqs)
	key := make(map[string]uint8, n)
	for i, r := range reqs {
		key[r] = uint8(i)
	}

	dp := make(map[int][]int, int(math.Pow(2, float64(n))))
	dp[0] = []int{}[:0:0]

	for i, person := range peoples {
		personSkills := 0
		for _, s := range person {
			if index, ok := key[s]; ok {
				personSkills |= 1 << index
			}
		}
		for teamSkills, team := range dp {
			m := len(team) + 1
			withPerson := teamSkills | personSkills
			if withPerson == teamSkills {
				continue
			}
			oldTeam, ok := dp[withPerson]
			if !ok || len(oldTeam) > m {
				dp[withPerson] = append(team, i)[:m:m]
			}
		}
	}

	return dp[1<<uint(n)-1]
}
