package problem0920

// ref: https://leetcode.com/problems/number-of-music-playlists/discuss/178415/C++JavaPython-DP-Solution

const mod = 1e9 + 7

func numMusicPlaylists(N int, L int, K int) int {
	dp := [101][101]int{}
	// dp[n][l] means n songs in l lines and K other songs has been played when some song wants to play again

	// 0 <= K < N <= L <= 100
	for n := K + 1; n <= N; n++ {
		// when n songs in n lines,
		//     just factorial n conditions
		dp[n][n] = factorial(n)
		// when n songs in n+ lines,
		for l := n + 1; l <= L; l++ {
			count := dp[n-1][l-1]*n + dp[n][l-1]*(n-K)
			// dp[n-1][l-1] *n
			//     dp[n-1][l-1] means n-1 songs has been in [1,l-1] lines,
			//     want to add a new song which is not in n-1 songs to l line
			//     the new song could be anyone of n songs, So *n
			// dp[n][l-1] *(n-K)
			//     dp[n][l-1] means n songs has been in [1,l-1] lines,
			//     add one of n songs to l line
			//     the song to add must be different with [l-k, l-1] line songs
			//     only choose from (n-k) songs, So *(n-k)
			dp[n][l] = count % mod
		}
	}
	return dp[N][L]
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return (n * factorial(n-1)) % mod
}
