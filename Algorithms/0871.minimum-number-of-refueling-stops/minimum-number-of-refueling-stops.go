package problem0871

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	size := len(stations)

	pq := make(PQ, 0, size*500)
	pq = append(pq, station{
		index: -1,
		miles: startFuel,
		stops: 0,
	})
	heap.Init(&pq)

	for len(pq) > 0 {
		s := heap.Pop(&pq).(station)
		if s.miles >= target {
			return s.stops
		}

		for i := s.index + 1; i < size && stations[i][0] <= s.miles; i++ {
			heap.Push(&pq, station{
				index: i,
				miles: s.miles + stations[i][1],
				stops: s.stops + 1,
			})
		}
	}

	return -1
}

// station 是 priorityQueue 中的元素
type station struct {
	stops, index, miles int
}

// PQ implements heap.Interface and holds entries.
type PQ []station

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].stops == pq[j].stops {
		// if pq[i].miles == pq[j].miles {
		// return pq[i].index > pq[j].index
		// }
		return pq[i].miles > pq[j].miles
	}
	return pq[i].stops < pq[j].stops
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(station)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
