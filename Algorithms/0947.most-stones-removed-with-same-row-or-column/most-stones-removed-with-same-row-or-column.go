package problem0947

// 如果把相同行或者列上的点，看作是联通的
// 那么本题的本质就是，每个联通区域只保留一个点，则可以删除多少个点

func removeStones(stones [][]int) int {
	u := newUnion(20000)
	for _, s := range stones {
		// 把一个点的 x 和 y 坐标进行连通，可以把复杂度从 O^2 降低到 O
		// 当为了避免坐标轴的重叠，需要把每个 y 坐标的值 +10000
		u.union(s[0], s[1]+10000)
	}

	keeps := make(map[int]int, 1000)
	for _, s := range stones {
		root := u.find(s[0])
		keeps[root]++
	}

	return len(stones) - len(keeps)
}

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
