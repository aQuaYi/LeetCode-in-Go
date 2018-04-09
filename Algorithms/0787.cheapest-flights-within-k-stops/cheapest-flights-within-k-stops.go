package problem0787

import "container/heap"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	fmap := make([][][]int, n)
	for i := range flights {
		src := flights[i][0]
		fmap[src] = append(fmap[src], flights[i][1:])
	}

	pq := make(PQ, 0, n)
	// 把 src 作为起点放入 pq
	heap.Push(&pq, &city{
		price:     0,     // 费用为 0
		id:        src,   // id 就是 src
		countdown: k + 1, // 飞机还可以飞行 k+1 次
	})

	for len(pq) > 0 {
		ct, _ := heap.Pop(&pq).(*city)

		// 如果到达了 dst
		if ct.id == dst {
			// 返回到达 ct 的票价即可
			// 优先队列保证了此时的 price 一定是最小的
			return ct.price
		}

		// 如果还能飞行
		if ct.countdown > 0 {
			nexts := fmap[ct.id]
			for _, n := range nexts {
				heap.Push(&pq, &city{
					id:        n[0],
					price:     ct.price + n[1],
					countdown: ct.countdown - 1,
				})
			}
		}
	}

	return -1
}

// city 是 priorityQueue 中的元素
type city struct {
	id, price, countdown int
	// countdown 表示飞机还可以飞行的次数
}

// PQ implements heap.Interface and holds entries.
type PQ []*city

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].price < pq[j].price
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 city
func (pq *PQ) Push(x interface{}) {
	temp := x.(*city)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 city
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
