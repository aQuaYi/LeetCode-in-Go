package problem0790

var mod = 1000000007
var results []int

func numTilings(N int) int {
	if results == nil {
		results = make([]int, 1001)
		results[0] = 1
		results[1] = 1
	}

	res := 0
	for i := 0; i < N; i++ {
		if results[i] == 0 {
			results[i] = numTilings(i)
		}
		res += results[i] * chain(N-i)
		res %= mod
	}

	return res
}

// chain 由砖组合而成，当沿着竖缝切下去的时候，总会破坏另一块砖
// chain(3) = 2 代表了 长度为 3 的 chain 一共有 2 块，它们是
//  XXY XYY
//  XYY XXY
// chain(5) = 2 代表了 长度为 5 的 chain 一共有 2 块，它们是
//  XXZZY XZZYY
//  XZZYY XXZZY
func chain(x int) int {
	if x >= 3 {
		return 2
	}
	return 1
}
