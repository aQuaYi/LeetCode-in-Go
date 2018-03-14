package problem0327

var lo, up int
var tmp []int

func countRangeSum(nums []int, lower int, upper int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	// sum[i] == sum(nums[:i])
	// sum[j] - sum[i] == sum(nums[i:j])
	sum := make([]int, size+1)
	for i, n := range nums {
		sum[i+1] = sum[i] + n
	}

	lo = lower
	up = upper
	tmp = make([]int, len(sum))

	return mergeSort(sum)
}

// 对前 i 项的和，进行归并排序
// 归并 left 和 right 的同时，检查符合条件的子数组
func mergeSort(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}

	mid := size / 2
	// left 中任意一个值，所代表的 sum(nums[:i])
	// right 中任意一个值，所代表的 sum(nums[:j])
	// 可以得知， i < j， 这一点很重要，
	// 正因为如此，对于任意一个 right[j]-left[i] 都是某个子数组的 sum，统计才有意义
	left := nums[:mid]
	right := nums[mid:]

	// 分而治之
	count := mergeSort(left) + mergeSort(right)

	// 用于归并
	i := 0
	j := 0

	// 用于 count
	cl := 0
	cr := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp[i+j] = left[i]

			for cl < len(right) && right[cl]-left[i] < lo {
				cl++
			}
			for cr < len(right) && right[cr]-left[i] <= up {
				cr++
			}
			// 对于 left[i] 所代表的 sum(nums[:x]) 来说
			// 此时 right 中存在有 cr - cl 个 right[cy] 所代表的 sum(nums[:y])
			// 能够保证 lower <= sum(nums[x:y]) <= upper
			//
			// 上段注释中，在 nums 中用到了新的索引号 x , y ，而不是 left 和 right中用到的 i 和 cy
			// 是因为 left 和 right 经过排序后，位置已经打乱，两者的索引号对应关系也被打乱了
			// 但是，我们知道 x < y 一定成立，就可以了
			count += cr - cl
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
		// left 中的元素没有检查完，继续检查
		for i < len(left) {
			for cl < len(right) && right[cl]-left[i] < lo {
				cl++
			}
			for cr < len(right) && right[cr]-left[i] <= up {
				cr++
			}
			count += cr - cl
			i++
		}
	}

	copy(nums, tmp)

	return count
}
