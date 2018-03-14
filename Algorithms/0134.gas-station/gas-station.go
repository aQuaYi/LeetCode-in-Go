package problem0134

func canCompleteCircuit(gas []int, cost []int) int {
	remains, debts, start := 0, 0, 0

	for i, g := range gas {
		remains += g - cost[i]
		if remains < 0 {
			// i + 1 处重新开始
			start = i + 1
			// 记录沿路一共欠缺的油量
			debts += remains
			// remain 至零
			remains = 0
		}
	}

	if debts+remains < 0 {
		// 最后的剩余的油量，如法全部偿还前期欠缺的油量
		// 则，无法跑完一圈
		return -1
	}

	return start
}
