package problem0830

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0, len(s)/3)
	l, r := 0, 1

	for ; r < len(s); r++ {
		if s[l] != s[r] {
			l = r
			continue
		}

		if r-l+1 == 3 { // 找到一个 large group
			res = append(res, []int{l, r})
		} else if r-l+1 > 3 { // 最后一个 large group 继续变大
			res[len(res)-1][1] = r
		}
	}

	return res
}
