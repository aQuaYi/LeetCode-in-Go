package problem0260

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}

	lowest := xor & -xor

	var a, b int
	for _, num := range nums {
		if num&lowest == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
