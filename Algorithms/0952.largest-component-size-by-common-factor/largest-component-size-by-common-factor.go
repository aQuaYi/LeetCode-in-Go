package problem0952

func largestComponentSize(A []int) int {
	size := len(A)
	u := newUnion(size)

	// 获取最大值
	max := -1
	for _, v := range A {
		if max < v {
			max = v
		}
	}

	halfMax := max / 2

	primes := []int{2, 3}
	for i := 5; i <= halfMax; i += 2 {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	for _, p := range primes {
		i := 0
		for i < size && A[i]%p != 0 {
			i++
		}

		for j := i + 1; j < size; j++ {
			if A[j]%p != 0 {
				continue
			}
			u.union(i, j)
		}
	}

	return u.max
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

type union struct {
	parents []int // 父链接数组(由触点索引)
	counts  []int // (由触点索引的) 各个根节点所对应的分量的大小
	max     int
}

func newUnion(N int) *union {
	parents := make([]int, N)
	for i := range parents {
		parents[i] = i
	}
	counts := make([]int, N)
	for i := range counts {
		counts[i] = 1
	}
	return &union{
		parents: parents,
		counts:  counts,
		max:     1,
	}
}

func (u *union) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *union) find(p int) int {
	// 跟随连接找到根节点
	if u.parents[p] != p {
		u.parents[p] = u.find(u.parents[p])
	}
	return u.parents[p]
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	if i == j {
		return
	}

	if u.counts[i] > u.counts[j] {
		i, j = j, i
	}

	// 将小树的根节点连接到大树的根节点
	u.parents[i] = j
	u.counts[j] += u.counts[i]
	u.max = max(u.max, u.counts[j])
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
