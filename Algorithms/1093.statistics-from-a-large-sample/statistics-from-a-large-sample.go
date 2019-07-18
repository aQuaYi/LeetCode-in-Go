package problem1093

func sampleStats(count []int) []float64 {
	return []float64{
		minimum(count),
		maximum(count),
		mean(count),
		median(count),
		mode(count),
	}
}

// minimum：最小值
func minimum(count []int) float64 {
	i := 0
	for ; i < len(count); i++ {
		if count[i] != 0 {
			break
		}
	}
	return float64(i)
}

// maximum：最大值
func maximum(count []int) float64 {
	i := len(count) - 1
	for ; i >= 0; i-- {
		if count[i] != 0 {
			break
		}
	}
	return float64(i)
}

// mean   ：平均数
func mean(count []int) float64 {
	sum, c := 0, 0
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		sum += i * count[i]
		c += count[i]
	}
	return float64(sum) / float64(c)
}

// median ：中位数，排序后，位于正中间的数。
func median(count []int) float64 {
	t := make([]int, 0, 256)
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		for j := 0; j < count[i]; j++ {
			t = append(t, i)
		}
	}
	n := len(t)
	if n%2 == 1 {
		return float64(t[n/2])
	}
	h := n / 2
	return float64(t[h-1]+t[h]) / 2
}

// mode   ：出现次数最多的数
func mode(count []int) float64 {
	k, c := -1, -1
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		if c < count[i] {
			k, c = i, count[i]
		}
	}
	return float64(k)
}
