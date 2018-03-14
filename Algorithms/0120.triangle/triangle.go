package problem0120

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// 从上往下，依次修改triangle
	// 使得 triangle[i][j] 表示到达 (i,j) 点的最小值
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			switch {
			case j == 0:
				triangle[i][0] += triangle[i-1][0]
			case j == i:
				triangle[i][i] += triangle[i-1][i-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	// 从最后一行，获取最小值
	min := triangle[n-1][0]
	for j := 1; j < n; j++ {
		if min > triangle[n-1][j] {
			min = triangle[n-1][j]
		}
	}

	return min
}
