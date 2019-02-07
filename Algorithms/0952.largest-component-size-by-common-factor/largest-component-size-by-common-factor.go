package problem0952

import "sort"

func largestComponentSize(A []int) int {
	size := len(A)
	u := newUnion(size)

	sort.Ints(A)

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			a, b := A[i], A[j]
			if u.isConnected(i, j) || gcd(a, b) > 1 {
				u.union(i, j)
			}
		}
	}

	return u.max
}

// Greatest Common Divisor(GCD)
func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

type union struct {
	id  []int // 父链接数组(由触点索引)
	sz  []int // (由触点索引的) 各个根节点所对应的分量的大小
	max int
}

func newUnion(N int) *union {
	id := make([]int, N)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, N)
	for i := range sz {
		sz[i] = 1
	}
	return &union{
		id:  id,
		sz:  sz,
		max: 1,
	}
}

func (u *union) isConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *union) find(p int) int {
	// 跟随连接找到根节点
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	if i == j {
		return
	}
	if u.sz[i] > u.sz[j] {
		i, j = j, i
	}
	// 将小树的根节点连接到大树的根节点
	u.id[i] = j
	u.sz[j] += u.sz[i]
	u.max = max(u.max, u.sz[j])
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
