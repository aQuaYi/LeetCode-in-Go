package unionfind

// leetcode 中需要 union 的集合都太小了，最普通的集合算法，就够用了

type union struct {
	parent []int
}

func newUnion(size int) *union {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &union{
		parent: parent,
	}
}

func (u *union) find(i int) int {
	if u.parent[i] == i {
		return i
	}
	u.parent[i] = u.find(u.parent[i])
	return u.parent[i]
}

func (u *union) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	u.parent[yp] = xp
}
