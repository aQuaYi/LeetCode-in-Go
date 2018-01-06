package Problem0576

const mod = 1e9 + 7

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func findPaths(m, n, N, i, j int) int {
	grid := [50][50]int{}
	grid[i][j] = 1

	queue := make([][]int, 1, m*n)
	queue[0] = []int{i, j}
	count := 1

	res := 0
	for {
		if count == 0 {
			count = len(queue)
			N--
			if N == 0 {
				break
			}
		}

		i, j = queue[0][0], queue[0][1]
		queue = queue[1:]
		count--

		if 0 <= i-1 {
			queue = append(queue, []int{i - 1, j})
		} else {
			res++
		}

		if 0 <= j-1 {
			queue = append(queue, []int{i, j - 1})
		} else {
			res++
		}

		if i+1 < m {
			queue = append(queue, []int{i + 1, j})
		} else {
			res++
		}

		if j+1 < n {
			queue = append(queue, []int{i, j + 1})
		} else {
			res++
		}

		res %= mod
	}

	return res
}
