package problem0773

import (
	"strings"
)

func slidingPuzzle(board [][]int) int {
	origin, target := "123450", convert(board)
	if origin == target {
		return 0
	}

	queue := make([]string, 1, 720)
	queue[0] = origin

	hasSeen := make(map[string]bool, 720)
	hasSeen[origin] = true

	res := 0
	d := []int{-1, 1, 3, -3}
	countDown := len(queue)
	for len(queue) > 0 {
		// 从队列中抽取 s 对其进行变形
		s := queue[0]
		i := strings.IndexByte(s, '0')
		for k := range d {
			j := i + d[k]
			if j < 0 ||
				j > 5 ||
				i == 2 && j == 3 ||
				i == 3 && j == 2 {
				continue
			}
			c := swap(s, i, j)
			if c == target {
				return res + 1
			}

			if !hasSeen[c] {
				queue = append(queue, c)
				hasSeen[c] = true
			}
		}

		// 删除队列头部，检查 res 是否需要 ++
		queue = queue[1:]
		countDown--
		if countDown == 0 {
			countDown = len(queue)
			res++
		}
	}

	return -1
}

func convert(board [][]int) string {
	res := make([]byte, 6)
	for i := range res {
		res[i] = byte(board[i/3][i%3]) + '0'
	}
	return string(res)
}

func swap(s string, i, j int) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]
	return string(b)
}
