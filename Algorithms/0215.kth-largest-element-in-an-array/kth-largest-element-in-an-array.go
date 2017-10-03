package Problem0215

func findKthLargest(nums []int, k int) int {
	heapify(nums)

	if k == 1 {
		return nums[0]
	}

	for i := 1; i < k; i++ {
		nums = nums[1:]
		fixDown(nums, 0)
	}

	return nums[0]
}

func heapify(a []int) {
	size := len(a)
	for i := size/2 - 1; i >= 0; i-- {
		fixDown(a, i)
	}
}

func fixDown(a []int, i int) {
	size := len(a)
	tmp := a[i]
	j := i*2 + 1
	if j+1 < size && a[j] < a[j+1] {
		j++
	}

	for j < size {
		if tmp > a[j] {
			break
		}

		a[i] = a[j]
		i = j
		j = j*2 + 1
		if j+1 < size && a[j] < a[j+1] {
			j++
		}
	}

	a[i] = tmp
}
