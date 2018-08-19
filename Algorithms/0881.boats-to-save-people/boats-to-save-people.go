package problem0881

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l, r := 0, len(people)-1

	res := 0

	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		res++
	}

	return res
}

// 首先看到 people.length <= 50000 就知道不可能是使用动态规划做。
// 然后，还要注意到一个关键条件：每艘船只能坐两个人。
// 可以使用贪心算法做
