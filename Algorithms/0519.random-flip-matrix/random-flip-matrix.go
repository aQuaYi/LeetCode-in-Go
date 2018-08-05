package problem0519

import "math/rand"

// Solution object will be instantiated and called as such:
// obj := Constructor(n_rows, n_cols);
// param_1 := obj.Flip();
// obj.Reset();
type Solution struct {
	rows, cols, total int
	recorder          map[int]int
}

// Constructor 构建 Solution
func Constructor(rows, cols int) Solution {
	r := make(map[int]int, rows)
	return Solution{
		rows:     rows,
		cols:     cols,
		total:    rows * cols,
		recorder: r,
	}
}

// Flip 选择 rows * cols 矩阵中的某个 0 进行翻转
func (s *Solution) Flip() []int {
	if s.total == 0 {
		panic("No Zero")
	}

	index := rand.Intn(s.total)
	cand := index
	if changed, ok := s.recorder[index]; ok {
		cand = changed
	}

	s.total--
	if changed, ok := s.recorder[s.total]; ok {
		s.recorder[index] = changed
	} else {
		s.recorder[index] = s.total
	}

	delete(s.recorder, s.total)

	r, c := cand/s.cols, cand%s.cols

	return []int{r, c}
}

// Reset 把 rows * cols 中的元素全部变成 0
func (s *Solution) Reset() {
	s.total = s.rows * s.cols
	s.recorder = make(map[int]int, s.rows)
}
