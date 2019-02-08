package problem0956

import "sort"

func tallestBillboard(rods []int) int {
	size := len(rods)
	sort.Ints(rods)

	total := 0
	for _, r := range rods {
		total += r
	}

	wastage := 0
	if total%2 == 1 {
		wastage = 1
	}

	for wastage < total {
		isUsed := make([]bool, size)
		sum := (total - wastage) / 2

		if rescur(wastage, 0, -1, rods, isUsed) &&
			rescur(sum, 0, -1, rods, isUsed) {
			return sum
		}

		wastage += 2
	}

	return 0
}

func rescur(goal, tmp, index int, rods []int, isUsed []bool) bool {
	if goal == tmp {
		return true
	}

	size := len(rods)

	if goal < tmp || index == size {
		return false
	}

	for i := index + 1; i < size; i++ {
		if goal < rods[i] {
			break
		}
		if isUsed[i] {
			continue
		}
		isUsed[i] = true
		if rescur(goal, tmp+rods[i], i, rods, isUsed) {
			return true
		}
		isUsed[i] = false
	}

	return false
}
