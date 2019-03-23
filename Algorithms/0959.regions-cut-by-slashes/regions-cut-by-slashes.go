package problem0959

func regionsBySlashes(grid []string) int {
	m := len(grid)
	size := m * m * 4
	u := newUnion(size)

	// cut every square to 4 parts, and mark them
	//    \0/
	//   3 X 1
	//    /2\
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			index := (i*m + j) * 4
			top := index + 0
			right := index + 1
			down := index + 2
			left := index + 3
			switch grid[i][j] {
			case '\\':
				u.union(top, right)
				u.union(down, left)
			case '/':
				u.union(top, left)
				u.union(down, right)
			default:
				u.union(top, right)
				u.union(right, down)
				u.union(down, left)
			}
			// union right square' left
			if rsl := index + 4 + 3; rsl < size {
				u.union(right, rsl)
			}
			// union down square' top
			if dst := index + 4*m; dst < size {
				u.union(down, dst)
			}
		}
	}

	return u.count
}

type union struct {
	parent []int
	count  int
}

func newUnion(size int) *union {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &union{
		parent: parent,
		count:  size,
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
	if xp != yp {
		u.parent[yp] = xp
		u.count--
	}
}
