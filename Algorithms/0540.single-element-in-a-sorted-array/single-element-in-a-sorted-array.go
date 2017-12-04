package Problem0540

func singleNonDuplicate(nums []int) int {
	n := len(nums)-1
	lo,hi:= 0 , n-1
	for lo< hi {
		mid := lo + (hi-lo)/2
if (mid %2==0 && (mid == n-1 || nums[mid]!= nums[mid+1]))||
(mid%2==1&&(mid==0|| nums[mid-1]!= nums[mid])) {
	return nums[mid]
}
	}
	res :=

	return res
}
