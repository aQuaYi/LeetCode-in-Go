package Problem0546

import "fmt"

func removeBoxes(boxes []int) int {
	m := make(map[string]int, len(boxes))

	var dfs func([]int) int
	dfs = func(boxes []int) int {
		key := fmt.Sprint(boxes)
		if v, ok := m[key]; ok {
			return v
		}

		if len(boxes) == 0 {
			m[key] = 0
			return 0
		}

		maxPoints := 0

		beg, end := 0, 0
		for beg < len(boxes) {
			for end < len(boxes) && boxes[beg] == boxes[end] {
				end++
			}
			maxPoints = max(maxPoints, dfs(append(boxes[:beg:beg], boxes[end:]...))+(end-beg)*(end-beg))
			beg = end
		}

		m[key] = maxPoints
		return maxPoints
	}

	return dfs(boxes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
