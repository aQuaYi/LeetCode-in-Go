package Problem0691

const intMax = 1<<63 - 1

func minStickers(stickers []string, target string) int {
	m := len(stickers)
	mp := make([][]int, m)
	for i := range mp {
		mp[i] = make([]int, 26)
		for _, c := range stickers[i] {
			mp[i][c-'a']++
		}
	}
	dp := make(map[string]int, m)
	dp[""] = 0

	return helper(dp, mp, target)
}

func helper(dp map[string]int, mp [][]int, target string) int {
	if count, ok := dp[target]; ok {
		return count
	}

	n := len(mp)

	tar := make([]int, 26)
	for _, c := range target {
		tar[c-'a']++
	}

	ans := 1<<63 - 1
	for i := 0; i < n; i++ {
		if mp[i][target[0]-'a'] == 0 {
			continue
		}

		s := make([]byte, 0, len(target))
		for j := 0; j < 26; j++ {
			if tar[j]-mp[i][j] > 0 {
				for k := tar[j] - mp[i][j]; 0 < k; k-- {
					s = append(s, byte('a'+j))
				}
			}
		}

		tmp := helper(dp, mp, string(s))
		if tmp != -1 {
			ans = min(ans, tmp+1)
		}
	}

	if ans == 1<<63-1 {
		dp[target] = -1
	} else {
		dp[target] = ans
	}

	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
