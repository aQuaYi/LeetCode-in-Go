package Problem0740

import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	cs := make(cards, 0, len(nums))
	c := card{num: nums[0], point: nums[0]}

	for i := 1; i < len(nums); i++ {
		if c.num == nums[i] {
			c.point += nums[i]
		} else {
			cs = append(cs, c)
			c = card{num: nums[i], point: nums[i]}
		}
	}
	cs = append(cs, c)

	res := 0
	sort.Sort(cs)

	isDeleted := make(map[int]bool, len(cs))
	for _, c := range cs {
		if isDeleted[c.num] {
			continue
		}
		res += c.point
		isDeleted[c.num+1] = true
		isDeleted[c.num-1] = true
	}

	return res
}

type card struct {
	num, point int
}

// cards 实现了 sort.Interface 接口
type cards []card

func (c cards) Len() int { return len(c) }

func (c cards) Less(i, j int) bool { return c[i].point > c[j].point }

func (c cards) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
