package problem0920

// ref: https://leetcode.com/problems/number-of-music-playlists/discuss/178415/C++JavaPython-DP-Solution

const mod = 1e9 + 7

func numMusicPlaylists(N int, L int, K int) int {
	dp := [101][101]int{}
	for i := K + 1; i <= N; i++ {
		for j := i; j <= L; j++ {
			if i == j || i == K+1 {
				dp[i][j] = factorial(i)
			} else {
				dp[i][j] = (dp[i-1][j-1]*i + dp[i][j-1]*(i-K)) % mod
			}
		}
	}
	return dp[N][L]
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
