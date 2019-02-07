package problem0952

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317}

func largestComponentSize(A []int) int {
	size := len(A)
	u := newUnion(size)

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
