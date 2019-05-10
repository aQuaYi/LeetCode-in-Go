package problem0997

import "sort"

func findJudge(N int, trust [][]int) int {
	if N == 1 {
		return 1
	}

	sort.Slice(trust, func(i int, j int) bool {
		return trust[i][1] < trust[j][1]
	})

	trustOther := [1001]bool{}
	for _, t := range trust {
		trustOther[t[0]] = true
	}

	for i := 0; i+N-2 < len(trust); i++ {
		judge := trust[i][1]
		if judge == trust[i+N-2][1] && !trustOther[judge] {
			return judge
		}
	}

	return -1
}
