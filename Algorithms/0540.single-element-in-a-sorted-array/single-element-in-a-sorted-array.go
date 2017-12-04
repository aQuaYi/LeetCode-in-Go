package Problem0540

func singleNonDuplicate(nums []int) int {
	n := len(nums)-1
	lo,hi:= 0 , n-1
	for lo< hi {
		mid := lo + (hi-lo)/2
		if (mid == 0&& nums[0]!= nums[1])||
		(mid == n-1 && nums[n-2]!=nums[n-1])||
		(nums[mid-1]!=)
	}
	res :=

	return res
}
