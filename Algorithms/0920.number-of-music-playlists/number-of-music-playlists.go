package problem0920

// ref: https://leetcode.com/problems/number-of-music-playlists/discuss/178415/C++JavaPython-DP-Solution

const mod = 1e9 + 7

func numMusicPlaylists(N int, L int, K int) int {
	dp := [101][101]int{}
	// dp[n][l] 就是题目要求的那个意思
	for n := K + 1; n <= N; n++ { // n <= K 没有意义，所以从 K+1 开始
		for l := n; l <= L; l++ { // 按照题设定 N<=L， 所以， n<=l
			if n == l || n == K+1 {
				// 此时，做一个 n 个元素的全排列即可，不会再有其他情况
				dp[n][l] = factorial(n)
			} else {
				dp[n][l] = (dp[n-1][l-1]*n + dp[n][l-1]*(n-K)) % mod
				// dp[n-1][l-1]*n : 如果 n-1 首歌已经符合要求地放入了前 l-1 个位置，那么第 n 首歌不会与 [0,l) 行中的歌重复，可以直接放入第 l 行。另外，第 n 首歌，可以是 n 首歌中的任意一首，所以，需要 *n。
				// dp[n][l-1]*(n-K) : 如果 n 首歌已经符合要求地放入了前 l-1 个位置，那么，放入第 l 行的歌不能与 [l-k,l) 行中的歌重复，只有 （n-k）首歌可以选，所以是 *(n-k)。
			}
		}
	}
	return dp[N][L]
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return (n * factorial(n-1)) % mod
	// * / % 具有相同的优先级，按照从左往右计算
}
