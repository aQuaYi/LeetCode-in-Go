package Problem0332

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	res := []string{}

	nexts := make(map[string][]string, len(tickets))

	for _, t := range tickets {
		nexts[t[0]] = append(nexts[t[0]], t[1])
	}

	for dep := range nexts {
		sort.Strings(nexts[dep])
	}

	return res
}

func lastDep(temp string) string {
	return temp[len(temp)-3 : len(temp)-1]
}
