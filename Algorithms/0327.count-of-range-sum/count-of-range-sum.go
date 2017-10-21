package Problem0327

var lo, hi int
var tmp []int

func countRangeSum(nums []int, lower int, upper int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	sum := make([]int, size+1)
	for i, n := range nums {
		sum[i+1] = sum[i] + n
	}

	lo = lower
	hi = upper
	tmp = make([]int, len(sum))

	return sort(sum)
}

func sort(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}

	mid := size / 2
	left := nums[:mid]
	right := nums[mid:]

	count := sort(left) + sort(right)

	i := 0
	j := 0

	sl := 0
	sr := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp[i+j] = left[i]
			for sl < len(right) && right[sl]-left[i] < lo {
				sl++
			}
			for sr < len(right) && right[sr]-left[i] <= hi {
				sr++
			}
			count += sr - sl
			i++
		} else {
			tmp[i+j] = right[j]
			j++
		}
	}

	if i == len(left) {
		copy(tmp[i+j:], right[j:])
	} else {
		copy(tmp[i+j:], left[i:])
		for i < len(left) {
			for sl < len(right) && right[sl]-left[i] < lo {
				sl++
			}
			for sr < len(right) && right[sr]-left[i] <= hi {
				sr++
			}
			count += sr - sl
			i++
		}
	}
	copy(nums, tmp)
	return count
}
