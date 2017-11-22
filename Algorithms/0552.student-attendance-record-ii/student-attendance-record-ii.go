package Problem0552

var m = 1000000007

func checkRecord(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 3
	case 2: // 后面初始化 A 的时候，会有 A[2]
		return 8
	}

	P := make([]int, n)
	P[0] = 1

	L := make([]int, n)
	L[0], L[1] = 1, 3

	A := make([]int, n)
	A[0], A[1], A[2] = 1, 2, 4

	for i := 1; i < n; i++ {
		P[i-1] %= m
		L[i-1] %= m
		A[i-1] %= m

		P[i] = A[i-1] + P[i-1] + L[i-1]

		if i > 1 {
			L[i] = A[i-1] + P[i-1] + A[i-2] + P[i-2]
		}

		if i > 2 {
			A[i] = A[i-1] + A[i-2] + A[i-3]
		}
	}

	return (P[n-1] + L[n-1] + A[n-1]) % m
}
