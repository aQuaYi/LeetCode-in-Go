package problem1202

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := newUnionFind(n)

	for _, p := range pairs {
		uf.connect(p[0], p[1])
	}

	groups := make(map[int][]int, n)
	for c, p := range uf.parent {
		p = uf.find(p)
		// 相连的索引值，全部放在一起
		groups[p] = append(groups[p], c)
	}

	bytes := []byte(s)
	res := make([]byte, n)
	for _, g := range groups {
		size := len(g)
		a := make([]int, size)
		copy(a, g)
		// a 中的索引值，按照其在 bytes 中对应的字符大小来排序
		sort.Slice(a, func(i, j int) bool {
			return bytes[a[i]] < bytes[a[j]]
		})
		// g 中的索引值，按照其自身的大小排序
		sort.Ints(g)
		// 越小的位置，放入的值也越小
		for i := 0; i < size; i++ {
			res[g[i]] = bytes[a[i]]
		}
	}

	return string(res)
}

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

func (uf *unionFind) connect(x, y int) {
	uf.parent[uf.find(x)] = uf.find(y)
}

func (uf *unionFind) find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.find(uf.parent[i])
	}
	return uf.parent[i]
}
