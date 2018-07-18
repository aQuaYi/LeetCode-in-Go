package problem0457

import (
	"unsafe"
)

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	// 缩小 nums[i] 的范围
	for i := 0; i < size; i++ {
		nums[i] %= size
	}

	// 获取 int 型变量的位数 - 1
	bits := uint(unsafe.Sizeof(size) - 1)

	for i, n := range nums {
		// 用于标记 nums[i]
		// 每个外层 for 循环的 mark 必须不一样，才能检查此次路径是否闭合
		// mark 还需要与 n 同符号
		mark := (i + size) * (n>>bits | 1)

		// 每次内层 for 循环只需要检查未检查过的节点即可。
		// 他们的特点是 -size < n && n != 0 && n < size
		for -size < n && n < size && n != 0 {
			// nums[i] 通过了 for 的检查条件
			// 现在标记 i 节点
			nums[i] = mark

			// 跳转到下一个节点
			// 更新 i 和 n
			i = (n + i + size) % size
			n = nums[i]

			if n == mark {
				// 发现闭环
				return true
			}

			if n*mark < 0 {
				// 出现转向
				break
			}
		}
	}

	return false
}
