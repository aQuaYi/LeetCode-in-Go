package problem0519

import "math/rand"

// Solution object will be instantiated and called as such:
// obj := Constructor(n_rows, n_cols);
// param_1 := obj.Flip();
// obj.Reset();
type Solution struct {
	rows, cols int
	total      int // 矩阵中剩余的 0 的个数
	rec        map[int]int
}

// Constructor 构建 Solution
func Constructor(rows, cols int) Solution {
	r := make(map[int]int)
	return Solution{
		rows:  rows,
		cols:  cols,
		total: rows * cols,
		rec:   r,
	}
}

// Flip 选择 rows * cols 矩阵中的某个 0 进行翻转
func (s *Solution) Flip() []int {
	if s.total == 0 {
		return nil
	}

	index := rand.Intn(s.total)

	cand := index
	if change, isFound := s.rec[index]; isFound {
		cand = change
	}
	r, c := cand/s.cols, cand%s.cols

	s.total--
	if change, isFound := s.rec[s.total]; isFound {
		s.rec[index] = change
	} else {
		s.rec[index] = s.total
	}

	return []int{r, c}
}

// Reset 把 rows * cols 中的元素全部变成 0
func (s *Solution) Reset() {
	s.total = s.rows * s.cols
	s.rec = make(map[int]int)
}
