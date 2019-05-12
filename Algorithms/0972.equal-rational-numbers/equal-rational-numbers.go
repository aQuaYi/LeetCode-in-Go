package problem0972

import (
	"strconv"
	"strings"
)

// 首先，牢记 note 中提到的 3 点
// 3. 1 <= <IntegerPart>.length <= 4
// 4. 0 <= <NonRepeatingPart>.length <= 4
// 5. 1 <= <RepeatingPart>.length <= 4
// 再根据 IEEE 754 中 https://zh.wikipedia.org/wiki/IEEE_754#%E8%AE%A8%E8%AE%BA%E4%B8%80 提到的 64 bit 双精度可以准确地表示 15 位十进制数
// 分两种情况讨论：
// 1： 没有重复部分，可以直接精准地转换成 float64
// 2： 存在重复部分，可以延展 S 到 20 位长度后，再近似地转换成 float64
//     因为 note 中规定了各个部分的长度，
//     A: 重复部分不为 9 的话，在 15 位以内就会出现差别
//     B: 重复部分为 9 的话， 至少 S[9:20] 中皆为 9，会在四舍五入时进位，与数学规定一致

func isRationalEqual(S string, T string) bool {
	return convert(S) == convert(T)
}

func convert(s string) float64 {
	i := strings.IndexByte(s, '(')
	if i != -1 {
		base, repeat := s[:i], s[i+1:len(s)-1]
		s = base + strings.Repeat(repeat, (20-len(base))/len(repeat)+1)
		s = s[:20] // 20 could be 19, but not 18
	}
	res, _ := strconv.ParseFloat(s, 64)
	return res
}
