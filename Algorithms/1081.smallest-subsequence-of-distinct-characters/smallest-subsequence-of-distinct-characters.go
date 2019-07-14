package problem1081

import "strings"

func smallestSubsequence(text string) string {
	n := len(text)

	last := [26]int{}
	for i, b := range text {
		last[b-'a'] = i
	}

	stack, top := make([]int, n), -1
	hasSeen := [26]bool{}
	for i := 0; i < n; i++ {
		c := int(text[i] - 'a')
		if hasSeen[c] {
			continue
		}
		for top >= 0 &&
			stack[top] > c &&
			i < last[stack[top]] {
			pop := stack[top]
			top--
			hasSeen[pop] = false
		}
		top++
		stack[top] = c
		hasSeen[c] = true
	}

	var sb strings.Builder
	for i := 0; i <= top; i++ {
		c := byte(stack[i] + 'a')
		sb.WriteByte(c)
	}
	return sb.String()
}
