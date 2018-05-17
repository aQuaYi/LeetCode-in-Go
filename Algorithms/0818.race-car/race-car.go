package problem0818

import (
	"math"
)

var dp [10001]int

func racecar(t int) int {
	if dp[t] > 0 {
		return dp[t]
	}

	// 2^(n-1) <= target < 2^n
	n := uint(math.Log2(float64(t))) + 1

	if t == 1<<n-1 {
		// n 个 A，就可以正好走到 t
		dp[t] = int(n)
	} else {
		// 一开始就 n 个 A，距离 (1<<n-1) > t 的情况
		dp[t] = racecar(1<<n-1-t) + int(n) + 1
		// 		^					^		 ^
		// 		|					|		 |
		// 		|					|		转向 R
		// 		|					n 个 A, 到达 1<<n-1 需要的步骤数目
		// 		走完剩下距离的最优解

		for m := uint(0); m < n-1; m++ {
			// 一开始就连续 (n-1) 个 A，到达 1<<(n-1)-1
			// 转向 R
			// 连续 m 个 A，到达 1<<(n-1) - 1 - (1<<m - 1) = 1<<(n-1) - 1<<m
			// 再一次转向 R，如果不转向的话，就是背朝着 target 了
			dp[t] = min(dp[t], racecar(t-1<<(n-1)+1<<m)+int(n-1)+1+int(m)+1)
			//					^						 ^	 	 ^  ^     ^
			// 					|						 |       |  |	  |
			// 					|						 |     	 |  |	转向 R
			// 					|	 					 | 	  	 |  m 个 A
			// 					| 						 |       转向 R
			// 					|						n-1 个 A
			// 					走完剩下的距离的最优解

		}
	}

	return dp[t]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
