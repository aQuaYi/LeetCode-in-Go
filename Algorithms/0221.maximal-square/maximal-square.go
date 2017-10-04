package Problem0221

func maximalSquare(matrix [][]byte) int {
	max := 0
	m := len(matrix)
	if m == 0 {
		return max
	}
	n := len(matrix[0])
	if n == 0 {
		return max
	}

	checked := make([][]bool, m)
	for i := range checked {
		checked[i] = make([]bool, n)
	}

	x := make([]int, 0, m*n*2)
	y := make([]int, 0, m*n*2)

	push := func(i, j int) {
		x = append(x, i)
		y = append(y, j)
	}

	pop := func() (i, j int) {
		i = x[0]
		x = x[1:]
		j = y[0]
		y = y[1:]
		return
	}

	isEmpty := func() bool {
		return len(x) == 0
	}

	clean := func() {
		x = x[len(x):]
		y = y[len(y):]
	}

	size := func() int {
		return len(x)
	}

	bfs := func(a, b int) {
		temp := 0
		push(a, b)
		var i, j, k int
		for !isEmpty() {
			k = size()
			for k > 0 {
				i, j = pop()
				k--
				if matrix[i][j] == '0' {
					clean()
					return
				}

				if i-a <= j-b {
					push(i, j+1)
				}
				if i-a == j-b {
					push(i+1, j+1)
				}
				if i-a >= j-b {
					push(i+1, j)
				}
			}
			temp++
			if max < temp {
				max = temp
			}

			if a+temp >= m || b+temp >= n {
				clean()
				return
			}
		}
	}

	for i := 0; i+max < m; i++ {
		for j := 0; j+max < n; j++ {
			bfs(i, j)
		}
	}

	return max * max
}
