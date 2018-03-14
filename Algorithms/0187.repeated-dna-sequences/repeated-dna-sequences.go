package problem0187

import (
	"sort"
)

func findRepeatedDnaSequences(s string) []string {
	var res []string
	if len(s) <= 10 {
		return nil
	}

	str := ""
	// rec 记录各个 subString 出现的次数
	rec := make(map[string]int, len(s)-9)
	for i := 0; i+10 <= len(s); i++ {
		str = s[i : i+10]
		if v := rec[str]; v == 1 {
			// 把已经出现过一次的 subString 存放入 res
			res = append(res, str)
		}
		rec[str]++
	}

	sort.Strings(res)

	return res
}
