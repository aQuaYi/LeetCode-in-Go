package problem0447

func numberOfBoomerangs(points [][]int) int {
	res := 0

	size := len(points)
	if size < 3 {
		return 0
	}

	for i := 0; i < size; i++ {
		// dMap[9]6 表示， points 中与 points[i] 距离为 3 的点，共有 6 个。
		dMap := make(map[int]int, size)

		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			d := dSquare(points[i], points[j])

			if _, ok := dMap[d]; ok {
				dMap[d]++
			} else {
				dMap[d] = 1
			}
		}

		for _, v := range dMap {
			// 相同距离的 v 个点
			// 总共有 v * (v - 1) 种排列
			res += v * (v - 1)
		}
	}

	return res
}

func dSquare(p1, p2 []int) int {
	x := p2[0] - p1[0]
	y := p2[1] - p1[1]
	return x*x + y*y
}
