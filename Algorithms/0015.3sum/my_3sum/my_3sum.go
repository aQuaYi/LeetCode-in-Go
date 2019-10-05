package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	/*
	:type: nums: []int
	:rtype: [][]int
	*/
	sort.Ints(nums)
	result := [][]int{}
	i := 0
	for i < len(nums) -2 {
		if (i == 0 || nums[i] != nums[i-1]) {
			j, k := i+1, len(nums) -1
			for j < k {
				s := nums[i] + nums[j] + nums[k]
				if s < 0 {
					j++
				} else if s > 0 {
					k--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					j, k = next(nums, j, k)
				}
			}
		}
		i++
	}
	return result
}

func next(nums []int, j int, k int) (int, int) {
	for j < k {
		switch {
			case nums[j] == nums[j+1]: 
				j++
			case nums[k] == nums[k-1]:
				k--
			default:
				j++
				k--
				return j, k
		}
	}
	return j, k
}
/*
func main() {
	res := threeSum([]int{-1,0,1,2,-1,-4})
	fmt.Println(res)
}
*/
