package problem0561

func arrayPairSum(nums []int) int {
	var buckets [20001]int8
	for _, num := range nums {
		buckets[num+10000]++
	}

	sum := 0
	needAdd := true
	for num, count := range buckets {
		for count > 0 {
			if needAdd {
				sum += num - 10000
			}
			needAdd = !needAdd
			count--
		}
	}

	return sum
}
