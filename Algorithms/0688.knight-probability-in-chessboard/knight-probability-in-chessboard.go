package problem0688

import (
	"math"
)

var dx = []int{-2, -2, -1, 1, 2, 2, 1, -1}
var dy = []int{-1, 1, -2, -2, 1, -1, 2, 2}

func knightProbability(N int, K int, r int, c int) float64 {
	// 为了快速创建 pre 变量，使用数组，而不是切片
	// pre[i][j] = k 表示
	// 在某一步的时候，棋盘上的 (i,j) 位置，会有 k 个棋子落在此位置上
	pre := [25][25]float64{}
	pre[r][c] = 1

	for k := 0; k < K; k++ {
		// 按照规则移动棋子后
		// next 记录了各个位置上可能落下的棋子数
		next := [25][25]float64{}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if pre[i][j] == 0 {
					continue
				}
				for m := 0; m < 8; m++ {
					x := dx[m] + i
					y := dy[m] + j
					if 0 <= x && x < N && 0 <= y && y < N {
						next[x][y] += pre[i][j]
					}
				}
			}
		}

		pre = next
	}

	count := 0.
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			count += pre[i][j]
		}
	}

	return float64(count) / math.Pow(8., float64(K))
}
