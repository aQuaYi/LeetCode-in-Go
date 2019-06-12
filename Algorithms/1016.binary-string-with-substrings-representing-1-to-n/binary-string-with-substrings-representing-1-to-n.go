package problem1016

import "strings"

func queryString(S string, N int) bool {
	end := N >> 1
	for N >= end {
		if !strings.Contains(S, binaryString(N)) {
			return false
		}
		N--
	}
	return true
}

func binaryString(n int) string {
	B := make([]byte, 0, 30)
	for n > 0 {
		B = append(B, byte(n&1)+'0')
		n >>= 1
	}
	swap(B)
	return string(B)
}

func swap(B []byte) {
	i, j := 0, len(B)-1
	for i < j {
		B[i], B[j] = B[j], B[i]
		i++
		j--
	}
}
