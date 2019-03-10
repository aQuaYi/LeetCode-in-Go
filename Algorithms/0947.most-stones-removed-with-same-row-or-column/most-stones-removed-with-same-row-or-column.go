package problem0947

func removeStones(stones [][]int) int {
	u := newUnion()

	for _, s := range stones {
		u.union(s[0], s[1]+1000)
	}

	mark := make(map[int]int)

	for _, s := range stones {
		parent := u.find(s[0])
		mark[parent]++
	}

	return len(stones) - len(mark)
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

// union is ...
type union struct {
	id   [20000]int // 父链接数组(由触点索引)
	size [20000]int // (由触点索引的) 各个根节点所对应的分量的大小
}

func newUnion() *union {
	id := [20000]int{}
	for i := range id {
		id[i] = i
	}
	sz := [20000]int{}
	for i := range sz {
		sz[i] = 1
	}
	return &union{
		id:   id,
		size: sz,
	}
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
	if u.size[i] > u.size[j] {
		i, j = j, i
	}
	// 将小树的根节点连接到大树的根节点
	u.id[i] = j
	u.size[j] += u.size[i]
	return
}
