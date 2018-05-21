package problem0825

func numFriendRequests(ages []int) int {
	count := [121]int{}
	for i := range ages {
		count[ages[i]]++
	}

	sums := [121]int{}
	sum := 0
	res := 0
	for a := range count {
		sum += count[a]
		sums[a] = sum

		if a <= 14 || count[a] == 0 {
			// 此时 a/2+7 >= a
			continue
		}

		b := a/2 + 7

		res += count[a] * (sums[a] - sums[b] - 1)
		// sums[a] - sums[b]，年龄段 (b,a] 中的人数
		// sums[a] - sums[b] - 1，减一是因为不能和自己交朋友
		//
		// 题目中的 3 条交友规则，归纳总结一下，就是
		// 年龄为 a 的人
		// 只能和年龄段 (a/2+7, a] 中除了自己以外的人交朋友，
	}

	return res
}

/**
 * 当 age[B] <= 0.5 * age[A] + 7 和 age[B] > age[A] 成立时
 * age[B] > 100 && age[A] < 100 一定成立，所以这个限制是障眼法
 *
**/
