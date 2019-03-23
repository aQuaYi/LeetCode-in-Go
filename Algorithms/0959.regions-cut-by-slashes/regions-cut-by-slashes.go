package problem0959

func regionsBySlashes(grid []string) int {
	m := len(grid)

	size := m * m * 4
	u := newUnion(size)

	// cut every square to 4 parts, and mark them
	//      \top/
	//       \0/
	// left 3 X 1 right
	//       /2\
	//     /down\
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			baseIndex := (i*m + j) * 4
			top := baseIndex + 0
			right := baseIndex + 1
			down := baseIndex + 2
			left := baseIndex + 3
			switch grid[i][j] {
			case '\\':
				u.unite(top, right)
				u.unite(down, left)
			case '/':
				u.unite(top, left)
				u.unite(down, right)
			default:
				u.unite(top, right)
				u.unite(right, down)
				u.unite(down, left)
			}
			// right part unites right square's left
			if j+1 < m {
				rsl := baseIndex + 4 + 3
				u.unite(right, rsl)
			}
			// down part unites down square's top
			if i+1 < m {
				dst := baseIndex + 4*m
				u.unite(down, dst)
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

func (u *union) unite(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
	u.count--
}
