package problem0952

func largestComponentSize(A []int) int {
	size := len(A)
	u := newUnion(size)

	// rec = map[factor]index
	rec := make(map[int]int, size*2)

	for i := 0; i < size; i++ {
		a := A[i]
		for f := 2; f*f <= a; f++ {
			if a%f != 0 {
				continue
			}
			if j, ok := rec[f]; ok {
				u.union(i, j)
			} else {
				rec[f] = i
			}
			d := a / f
			if j, ok := rec[d]; ok {
				u.union(i, j)
			} else {
				rec[d] = i
			}
		}
		// a 本身还可以作为 factor
		if j, ok := rec[a]; ok {
			u.union(i, j)
		} else {
			rec[a] = i
		}
	}

	return u.max
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

type union struct {
	parents []int // 父链接数组(由触点索引)
	sizes   []int // (由触点索引的) 各个根节点所对应的分量的大小
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
		sizes:   counts,
		max:     1,
	}
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

	if u.sizes[i] > u.sizes[j] {
		i, j = j, i
	}

	// 将小树的根节点连接到大树的根节点
	u.parents[i] = j
	u.sizes[j] += u.sizes[i]
	u.max = max(u.max, u.sizes[j])
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
