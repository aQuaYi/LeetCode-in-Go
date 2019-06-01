package problem0990

func equationsPossible(equations []string) bool {
	u := newUnion(26)

	for _, e := range equations {
		if e[1] == '=' {
			u.unite(e[0]-'a', e[3]-'a')
		}
	}

	for _, e := range equations {
		if e[1] == '!' && u.find(e[0]-'a') == u.find(e[3]-'a') {
			return false
		}
	}

	return true
}

type union struct {
	parent []byte
}

func newUnion(size int) *union {
	parent := make([]byte, size)
	for i := range parent {
		parent[i] = byte(i)
	}
	return &union{
		parent: parent,
	}
}

func (u *union) find(i byte) byte {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *union) unite(x, y byte) {
	u.parent[u.find(x)] = u.find(y)
}
