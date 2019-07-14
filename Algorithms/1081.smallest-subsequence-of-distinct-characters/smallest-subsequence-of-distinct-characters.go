package problem1081

import "strings"

func smallestSubsequence(S string) string {
	n := len(S)

	last := [26]int{}
	for i, c := range S {
		last[c-'a'] = i
	}

	stack, top := make([]int, n), -1
	hasSeen := [26]bool{}
	for i := 0; i < n; i++ {
		c := int(S[i] - 'a')
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
		b := byte(stack[i] + 'a')
		sb.WriteByte(b)
	}

	return sb.String()
}
