package problem0049

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := make(map[string][]string)

	for _, str := range strs {
		temp := sortString(str)
		record[temp] = append(record[temp], str)
	}
	for _, v := range record {
		sort.Strings(v)
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	bytes := []byte(s)

	temp := make([]int, len(bytes))
	for i, b := range bytes {
		temp[i] = int(b)
	}

	sort.Ints(temp)

	for i, v := range temp {
		bytes[i] = byte(v)
	}

	return string(bytes)
}
