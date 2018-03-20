package problem0781

func numRabbits(answers []int) int {
	count := [1000]int{}
	for _, v := range answers {
		count[v]++
	}

	res := 0

	// 假设所有的 answer 都是 3
	// answer  count -> res
	//   3       1       4
	//   3       2       4
	//   3       3       4
	//   3       4       4
	//   3       5       8
	//   3       6       8
	//   3       7       8
	//   3       8       8
	//   3       9      12

	for ans, c := range count {
		if c == 0 {
			continue
		}
		ans++
		res += c / ans * ans
		if c%ans > 0 {
			res += ans
		}
	}

	return res
}
