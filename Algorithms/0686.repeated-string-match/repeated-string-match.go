package problem0686

func repeatedStringMatch(A string, B string) int {
	// ref: https://leetcode.com/problems/repeated-string-match/discuss/108084/C%2B%2B-4-lines-O(m-*-n)-or-O(1)-and-KMP-O(m-%2B-n)-or-O(n)
	n, m := len(A), len(B)
	prefTable := make([]int, m+1)
	for sp, pp := 1, 0; sp < m; {
		if B[pp] == B[sp] {
			sp++
			pp++
			prefTable[sp] = pp
		} else if pp == 0 {
			sp++
			prefTable[sp] = pp
		} else {
			pp = prefTable[pp]
		}
	}

	for i, j := 0, 0; i < n; i, j = i+max(1, j-prefTable[j]), prefTable[j] {
		for j < m && A[(i+j)%n] == B[j] {
			j++
		}
		if j == m {
			if (i+j)%n == 0 {
				return (i + j) / n
			} else {
				return (i+j)/n + 1
			}
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
