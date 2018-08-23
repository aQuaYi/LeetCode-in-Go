package problem0881

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	thin, fat := 0, len(people)-1

	res := 0

	for thin <= fat {
		// 如果，队列中最瘦的人，也可以上船的话，就跟上去
		if people[thin]+people[fat] <= limit {
			thin++
		}
		// 队列中，最胖的人肯定要上船
		fat--
		res++
	}

	return res
}

// 首先看到 people.length <= 50000 就应该知道不可能使用动态规划。
// 然后，还要注意到一个关键条件：每艘船只能坐两个人。
// 可以使用贪心算法做
