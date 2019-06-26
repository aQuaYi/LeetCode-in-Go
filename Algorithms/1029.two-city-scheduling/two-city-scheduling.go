package problem1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i int, j int) bool {
		return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0]
		// person i is either to A or to B
		// c = costs[i][1]-costs[i][0] means
		// if person i go to B, the cost have to be paid more.
		// so, c is larger, person i is more should go to A.
		// after sorting
		// persion[:N] go to A, the others go to B.
	})
	N := len(costs) / 2
	sum := 0
	for i := 0; i < N; i++ {
		sum += costs[i][0] + costs[i+N][1]
	}
	return sum
}
