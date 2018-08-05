package problem0519

import "math/rand"

// Solution object will be instantiated and called as such:
// obj := Constructor(n_rows, n_cols);
// param_1 := obj.Flip();
// obj.Reset();
type Solution struct {
	rows, cols, total int
	recorder          []int
}

// Constructor 构建 Solution
func Constructor(rows, cols int) Solution {
	r := make([]int, rows*cols)
	for i := range r {
		r[i] = i
	}
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
	cand := s.recorder[index]
	r, c := cand/s.cols, cand%s.cols

	s.total--
	s.recorder[index] = s.recorder[s.total]

	return []int{r, c}
}

// Reset 把 rows * cols 中的元素全部变成 0
func (s *Solution) Reset() {
	s.total = s.rows * s.cols

	r := make([]int, s.total)
	for i := range r {
		r[i] = i
	}
	s.recorder = r
}
