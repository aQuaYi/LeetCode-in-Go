package problem0452

import "sort"

func findMinArrowShots(bs [][]int) int {
	size := len(bs)
	if size == 0 {
		return 0
	}

	sort.Sort(balloons(bs))

	res := 0
	end := bs[0][1]

	for i := 1; i < size; i++ {
		if bs[i][0] <= end {
			continue
		}
		res++
		end = bs[i][1]
	}

	res++

	return res
}

// balloons 实现了 sort.Interface 接口
type balloons [][]int

func (b balloons) Len() int { return len(b) }

// 以 end 的大小来排序
func (b balloons) Less(i, j int) bool { return b[i][1] < b[j][1] }

func (b balloons) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
