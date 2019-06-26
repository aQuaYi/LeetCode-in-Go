package problem1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i int, j int) bool {
		if costs[i][0] == costs[j][0] {
			return costs[i][1] > costs[j][1]
		}
		return costs[i][0] < costs[j][0]
	})
	N := len(costs) / 2
	sum := 0
	for i := 0; i < N; i++ {
		sum += costs[i][0] + costs[i+N][1]
	}
	return sum
}
