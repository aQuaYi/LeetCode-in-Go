package problem0451

import "sort"

func frequencySort(s string) string {
	// r 记录了各个字符出现的次数
	r := ['z' + 1]int{}
	for i := range s {
		r[s[i]]++
	}

	// ss 中是由相同字符组成的 string 片段
	ss := make([]string, 0, len(s))
	for i := range r {
		if r[i] == 0 {
			continue
		}
		ss = append(ss, makeString(byte(i), r[i]))
	}

	// 按照题目中的规则，对 ss 进行排序
	sort.Sort(segments(ss))

	// 按顺序把 ss 中的片段拼接成 答案
	res := ""
	for _, s := range ss {
		res += s
	}

	return res
}

// 返回由 n 个 b 组成的字符串
func makeString(b byte, n int) string {
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = b
	}
	return string(bytes)
}

// segments 实现了 sort.Interface 接口
type segments []string

func (s segments) Len() int { return len(s) }

func (s segments) Less(i, j int) bool {
	// 按照题意，这个 if 不是必须的
	// 但是，添加以后，可以保证答案的唯一性
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s segments) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
