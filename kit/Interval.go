package kit

// Interval 提供区间表示
type Interval struct {
	Start int
	End   int
}

// Interval2Ints 把 Interval 转换成 整型切片
func Interval2Ints(i Interval) []int {
	return []int{i.Start, i.End}
}

// IntervalSlice2Intss 把 []Interval 转换成 [][]int
func IntervalSlice2Intss(is []Interval) [][]int {
	res := make([][]int, 0, len(is))
	for i := range is {
		res = append(res, Interval2Ints(is[i]))
	}
	return res
}
