package Problem0301

import "sort"

func removeInvalidParentheses(s string) []string {
	res := []string{}
	var dfs func(string, int, int)
	dfs = func(s string, l, r int) {
		if l+r == 0 && isOK(s) {
			res = append(res, s)
			return
		}

		last, temp := s, ""
		if r > 0 {
			for i := range s {
				if s[i] == ')' {
					temp = s[:i] + s[i+1:]
					if temp == last {
						continue
					}

					last = temp
					dfs(temp, l, r-1)
				}
			}
		} else if l > 0 {
			for i := len(s) - 1; i >= 0; i-- {
				if s[i] == '(' {
					temp = s[:i] + s[i+1:]
					if temp == last {
						continue
					}
					last = temp
					dfs(temp, l-1, r)
				}
			}
		}
	}

	l, r := count(s)
	dfs(s, l, r)

	sort.Strings(res)
	temp := make([]string, 1, len(res))
	temp[0] = res[0]
	for i := 1; i < len(res); i++ {
		if res[i] != res[i-1] {
			temp = append(temp, res[i])
		}
	}

	return temp
}

func count(s string) (left, right int) {
	for i := range s {
		switch s[i] {
		case '(':
			left++
		case ')':
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}

	return
}

func isOK(s string) bool {
	l, r := count(s)
	return l+r == 0
}
