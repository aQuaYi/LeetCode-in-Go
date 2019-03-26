package problem0964

import "math"

func leastOpsExpressTarget(x int, target int) int {
	if target == x {
		return 0
	} else if target == 1 {
		return 1
	}

	root := math.Log10(float64(target)) / math.Log10(float64(x))

	base := int(math.Pow(float64(x), math.Floor(root)))
	intRoot := int(math.Floor(root))
	if base == target {
		return intRoot - 1
	}

	return min(
		intRoot+leastOpsExpressTarget(x, abs(base-target)),
		intRoot+1+leastOpsExpressTarget(x, abs(base*x-target)),
	)

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
