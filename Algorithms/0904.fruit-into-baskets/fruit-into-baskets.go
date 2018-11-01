package problem0904

// 找到最长的 tree 的子数组，要求最多只能含有两种数字
func totalFruit(tree []int) int {
	tailType, theOther := -1, -1
	count, countTail := 0, 0 // 子数组的长度，子数组尾部相同类型水果的长度
	// NOTICE: countTail 不是子数组中全部 tailType 的个数，仅仅是尾部的那一部分
	res := 0

	for _, fruit := range tree {
		count++
		if fruit != tailType && fruit != theOther {
			// fruit 是新出现的水果类型
			count = countTail + 1
		}

		res = max(res, count)

		countTail++
		if fruit != tailType {
			// 子数组的尾部变成了新的类型
			theOther, tailType = tailType, fruit
			countTail = 1 // fruit 只有一个
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ref: https://leetcode.com/problems/fruit-into-baskets/discuss/170745/Problem:-Longest-Subarray-With-2-Elements
// Explanation:
// Loop all fruit c in tree,
// Note that a and b are the last two different types of fruit that we met,
// c is the current fruit type,
// so it's something like "....aaabbbc..."

// Case 1 c == b:
// fruit c already in the basket,
// and it's same as the last type of fruit
// cur += 1
// count_b += 1

// Case 2 c == a:
// fruit c already in the basket,
// but it's not same as the last type of fruit
// cur += 1
// count_b = 1
// a = b, b = c

// Case 3 c != b && c!= a:
// fruit c not in the basket,
// cur = count_b + 1
// count_b = 1
// a = b, b = c

// Of course, in each turn we need to update res = max(res, cur)

// Complexity:
// O(N) time, O(1) space
