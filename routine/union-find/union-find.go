package unionfind

// leetcode 中需要 union 的集合都太小了，最普通的集合算法，就够用了

type unionFind struct {
	parent []int
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{
		parent: parent,
	}
}

func (u *unionFind) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *unionFind) union(x, y int) {
	u.parent[u.find(x)] = u.find(y)
}
