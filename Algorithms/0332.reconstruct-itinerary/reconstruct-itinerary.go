package Problem0332

import "strings"

func findItinerary(tickets [][]string) []string {
	res := string('Z' + 1)

	var dfs func(string, [][]string)
	dfs = func(departs string, tickets [][]string) {
		if departs > res {
			return
		}

		if len(tickets) == 0 && res > departs {
			res = departs
		}

		for i, t := range tickets {
			if lastDep(departs) == t[0] {
				dfs(departs+" "+t[1], append(tickets[:i:i], tickets[i+1:]...))
			}
		}
	}

	for i, t := range tickets {
		if t[0] == "JFK" {
			dfs("JFK "+t[1], append(tickets[:i:i], tickets[i+1:]...))
		}
	}

	return strings.Split(res, " ")
}

func lastDep(temp string) string {
	return temp[len(temp)-3:]
}
