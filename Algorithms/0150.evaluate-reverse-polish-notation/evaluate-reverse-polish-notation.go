package problem0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// 用于存放数字的栈
	nums := make([]int, 0, len(tokens))
	for _, s := range tokens {
		if s == "+" ||
			s == "-" ||
			s == "*" ||
			s == "/" {
			// 遇到操作符， 数字出栈
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			// 运算后的结果，重新入栈
			nums = append(nums, compute(a, b, s))
		} else {
			// 遇到数字，则直接入栈
			temp, _ := strconv.Atoi(s)
			nums = append(nums, temp)
		}
	}

	return nums[0]
}

// 计算
func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
