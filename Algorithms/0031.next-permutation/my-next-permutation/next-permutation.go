package problem0031
import (
	"fmt"
)


func nextPermutation(nums []int) {
	left := len(nums) - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	fmt.Println(nums)
	// nums[left+1:] is currnetly an decreasing sequence
	reverse(nums, left+1)
	fmt.Println(nums)
	// nums[left+1:] is now an increasing sequence
	if left == -1 {
		return
	}
	//right := search(nums, nums[left])
	right := search1(nums, left+1, nums[left])
	// switch nums[left] with nums[right]
	nums[left], nums[right] = nums[right], nums[left]
}

func reverse(seq []int, l int) {
	for i, j := l, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
}
/*
seach and return the maximum value among nums[left+1:] that is greater than nums[left]
nums[left+1] is an increasing sequence
*/
func search(seq []int,target int) int {
	right := len(seq) -1
	for right >= 0 && seq[right] > target {
		right--
	}
	return right+1
}
/*
optimize search using bisect search
*/
func search1(seq []int, left int, target int) int {
	right := len(seq) - 1
	left--
	for left + 1 < right {
		mid := (left + right) / 2
		if seq[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
