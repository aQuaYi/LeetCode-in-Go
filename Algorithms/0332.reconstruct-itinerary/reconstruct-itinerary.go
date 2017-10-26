package Problem0332

import "sort"

func findItinerary(tickets [][]string) []string {

	nexts := make(map[string][]string, len(tickets))

	for _, t := range tickets {
		nexts[t[0]] = append(nexts[t[0]], t[1])
	}

	for k := range tickets {
		sort.Strings(tickets[k])
	}

	size := len(tickets) + 1
	res := make([]string, 1, size)
	res[0] = "JFK"

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == size {
			return true
		}

		old := nexts[res[idx-1]]
		size := len(old)
		var i int

		for i = 0; i < size; i++ {
			res[idx] = old[i]
			nexts[res[idx-1]] = append(old[:i:i], old[i+1:]...)
			if dfs(idx + 1) {
				return true
			}
		}

		return false
	}
	dfs(1)

	return res
}
