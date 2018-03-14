package problem0032

func longestValidParentheses(s string) int {
	var left, max, temp int
	record := make([]int, len(s))

	// 统计Record
	for i, b := range s {
		if b == '(' {
			left++
		} else if left > 0 {
			left--
			record[i] = 2
		}
	}

	// 修改record
	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
	}

	// 统计结果
	for _, r := range record {
		if r == 0 {
			temp = 0
			continue
		}

		temp++
		if temp > max {
			max = temp
		}
	}

	return max
}
