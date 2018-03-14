package problem0773

import (
	"strings"
)

func slidingPuzzle(board [][]int) int {
	hasSeen := make(map[string]bool, 720)
	q := make([]string, 1, 720)
	q[0] = convert(board)
	hasSeen[q[0]] = true
	target := "123450"

	res := 0

	if q[0] == target {
		return res
	}

	d := []int{-1, 1, 3, -3}
	size := len(q)
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		size--
		if size == 0 {
			size = len(q)
			res++
		}
		i := strings.IndexByte(s, '0')
		for k := range d {
			j := i + d[k]
			if j < 0 ||
				j > 5 ||
				i == 2 && j == 3 ||
				i == 3 && j == 2 {
				continue
			}
			b := []byte(s)
			b[i], b[j] = b[j], b[i]

			c := string(b)
			if c == target {
				return res
			}

			if !hasSeen[c] {
				q = append(q, c)
				hasSeen[c] = true
			}
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
