package problem0773

import (
	"strings"
)

func slidingPuzzle(board [][]int) int {
	original, target := "123450", convert(board)
	if original == target {
		return 0
	}

	q := make([]string, 1, 720)
	q[0] = original

	hasSeen := make(map[string]bool, 720)
	hasSeen[original] = true

	res := 0 // q 中最后一个候选项是 origin 经过 res 次变化得来的
	d := []int{-1, 1, 3, -3}
	size := len(q)
	for len(q) > 0 {
		//
		s := q[0]
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
				q = append(q, c)
				hasSeen[c] = true
			}
		}

		q = q[1:]
		size--
		if size == 0 {
			size = len(q)
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
