package problem0761

import (
	"sort"
	"strings"
)

func makeLargestSpecial(S string) string {
	if len(S) == 2 || len(S) == 0 {
		// 此时，Ｓ 必定为 Largest Special
		return S
	}

	// 把 S 尽可能多的划分成小的 special binary
	// 把划分成出的片段放入 res 中
	res := []string{}
	// count 用于在划分时，统计片段中 1 和 0 的个数是否相等
	count, i := 0, 0

	for j, b := range S {
		if b == '1' {
			count++
		} else {
			count--
		}

		if count == 0 {
			// 此时 S[i:j+1] 是 special binary
			// 且根据 special binary 的定义，可知
			// S[i] == '1'
			// S[j] == '0'
			// 另外
			// 不把　"1"+makeLargestSpecial(S[i+1:j])+"0"　写成
			// makeLargestSpecial(S[i:j+1]) 是因为
			// 对于　"111000" 来说
			// i == 0 , j+1 == len(S)
			// 递归的子集并没有缩小，程序无法结束
			res = append(res, "1"+makeLargestSpecial(S[i+1:j])+"0")
			i = j + 1
		}
	}

	// 按照从大到小的顺序排列 res 中的字符串
	sort.Sort(sort.Reverse(sort.StringSlice(res)))

	return strings.Join(res, "")
}
