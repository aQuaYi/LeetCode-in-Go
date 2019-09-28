package problem1202

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	un := newUnion(n)

	for _, p := range pairs {
		un.unite(p[0], p[1])
	}

	group := make(map[int][]int, n)
	for c, p := range un.parent {
		group[un.find(p)] = append(group[un.find(p)], c)
	}

	bytes := []byte(s)
	res := make([]byte, n)
	for _, children := range group {
		size := len(children)
		t := make([]int, size)
		copy(t, children)
		if size > 1 {
			sort.Slice(t, func(i, j int) bool {
				return bytes[t[i]] < bytes[t[j]]
			})
			sort.Ints(children)
		}
		for i := 0; i < size; i++ {
			res[children[i]] = bytes[t[i]]
		}
	}

	return string(res)
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

func (u *union) unite(x, y int) {
	u.parent[u.find(x)] = u.find(y)
}

func (u *union) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}
