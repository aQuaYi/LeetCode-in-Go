package problem0947

func removeStones(stones [][]int) int {
	u := newUnion(20000)

	for _, s := range stones {
		u.union(s[0], s[1]+10000)
	}

	roots := make(map[int]int, 1000)

	for _, s := range stones {
		root := u.find(s[0])
		roots[root]++
	}

	return len(stones) - len(roots)
}

type union struct {
	parent []int // 父链接数组(由触点索引)
}

func newUnion(size int) *union {
	p := make([]int, size)
	for i := range p {
		p[i] = i
	}
	return &union{
		parent: p,
	}
}

func (u *union) find(i int) int {
	if u.parent[i] == i {
		return i
	}
	u.parent[i] = u.find(u.parent[i])
	return u.parent[i]
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	u.parent[j] = i
}
