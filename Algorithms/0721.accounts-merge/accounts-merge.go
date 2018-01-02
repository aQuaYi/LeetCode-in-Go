package Problem0721

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	isCheckedEmail := make(map[string]int, len(accounts)*3)
	for i := 1; i < len(accounts[0]); i++ {
		isCheckedEmail[accounts[0][i]] = 0
	}

	res := make([][]string, 1, len(accounts))
	res[0] = accounts[0]

	for i := 1; i < len(accounts); i++ {
		same := -1

		for j := 1; j < len(accounts[i]); j++ {
			if id, ok := isCheckedEmail[accounts[i][j]]; ok {
				same = id
				break
			}
		}

		if same == -1 {
			res = append(res, accounts[i])
			id := len(res) - 1
			for j := 1; j < len(accounts[i]); j++ {
				isCheckedEmail[accounts[i][j]] = id
			}
		} else {
			for j := 1; j < len(accounts[i]); j++ {
				if _, ok := isCheckedEmail[accounts[i][j]]; !ok {
					isCheckedEmail[accounts[i][j]] = same
					res[same] = append(res[same], accounts[i][j])
				}
			}
			sort.Strings(res[same])
		}

	}

	return res
}
