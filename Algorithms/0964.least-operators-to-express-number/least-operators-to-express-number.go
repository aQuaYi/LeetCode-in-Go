package problem0964

import (
	"fmt"
	"math"
)

func leastOpsExpressTarget(x int, target int) int {
	fmt.Println("-=-=-=-=-=-=-=-=-")
	res := math.MaxInt64
	helper(x, target, 0, &res)
	return res
}

func helper(x, target, count int, res *int) {
	fmt.Println(target, count)
	if count >= *res {
		return
	}

	if target == x {
		*res = min(*res, count)
		fmt.Println("-", *res)
		return
	} else if target == 1 {
		*res = min(*res, count+1)
		fmt.Println("-", *res)
		return
	}

	root := math.Log10(float64(target)) / math.Log10(float64(x))

	base := int(math.Pow(float64(x), math.Floor(root)))
	intRoot := int(math.Floor(root))
	if base == target {
		*res = min(*res, count+intRoot-1)
		fmt.Println("-", *res)
		return
	}

	intRoot = max(intRoot, 1)
	if target-base <= base*x-target {
		if base == 1 {
			helper(x, target-base, count+intRoot+1, res)
		} else {
			helper(x, target-base, count+intRoot, res)
		}
		helper(x, base*x-target, count+intRoot+1, res)
	} else {
		if base == 1 {
			helper(x, base*x-target, count+intRoot, res)
		} else {
			helper(x, base*x-target, count+intRoot+1, res)
		}
		helper(x, target-base, count+intRoot, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
