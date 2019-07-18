package problem1093

import "sort"

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
	k := make([]int, 0, 256)
	c := make([]int, 0, 256)
	sum := 0
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		sum += count[i]
		k = append(k, i)
		c = append(c, sum)
	}
	if sum%2 == 1 {
		h := sum/2 + 1
		i := sort.SearchInts(c, h)
		return float64(k[i])
	}
	h := sum / 2
	i := sort.SearchInts(c, h)
	j := sort.SearchInts(c, h+1)
	return float64(k[i]+k[j]) / 2
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
