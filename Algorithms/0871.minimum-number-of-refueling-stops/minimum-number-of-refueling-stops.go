package problem0871

import "container/heap"

// 把题目换一种说法，就好理解了。
// 汽车在开往目的地的过程中，
// 会在沿路的加油站，都买上一箱汽油，
// 每个加油站的汽油大小还不一样。
// 汽车每次没油的时候，就在买过的汽油中，挑一箱加上。
// 问，汽车达到目的地的时候，最少需要加几次油？

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	size := len(stations)
	// gases 内放入一箱体积为 0 的汽油，是为了让 for 循环跑起来
	gases := make(intHeap, 1, size+1)
	miles := startFuel
	stops := 0
	i := 0

	for len(gases) > 0 {
		if miles >= target {
			// 到达了目的地
			return stops
		}

		// 路过加油站的时候，买汽油
		for i < size && stations[i][0] <= miles {
			heap.Push(&gases, stations[i][1])
			i++
		}

		// 在一堆汽油中，挑一箱最大的
		maxGas := heap.Pop(&gases).(int)
		stops++
		// 加上油，继续跑
		miles += maxGas
	}

	return -1
}

// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	// 返回堆中的最大值
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
