package problem0421

func findMaximumXOR(nums []int) int {
	var max, mask int

	// Basically we use the fact that Num1 ^ Num2 = Result ==> Num1 ^ Result = Num2
	// We start with the highest bit since we try to find the max to build the possible
	// result. All numbers with higher bits (have been processd or to be processed) are
	// hashed, then we use Num1 ^ Result to look up hash, if Num2 exists, then it is a
	// valid result, save it and go on to the next bit.
	for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)

		nMap := make(map[int]bool)
		for _, num := range nums {
			nMap[num&mask] = true
		}

		tmp := max | (1 << uint32(i))
		for key := range nMap {
			if nMap[tmp^key] {
				max = tmp
				break
			}
		}
	}

	return max
}
