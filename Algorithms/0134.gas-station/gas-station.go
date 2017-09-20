package Problem0134

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	if n == 0 {
		return -1
	}

	gas = append(gas, gas...)
	cost = append(cost, cost...)

	for i := 0; i < n; i++ {
		if success(gas[i:i+n], cost[i:i+n]) {
			return i
		}
	}

	return -1
}

func success(gas, cost []int) bool {
	gSum := 0
	cSum := 0

	for i := 0; i < len(gas); i++ {
		gSum += gas[i]
		cSum += cost[i]

		if gSum < cSum {
			return false
		}
	}

	return true
}
