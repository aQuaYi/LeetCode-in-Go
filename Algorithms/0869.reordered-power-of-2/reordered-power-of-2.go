package problem0869

import (
	"fmt"
	"strings"
)

// 对 [1, 1e9] 中所有的 2^n 进行编码
var isVaild = map[string]bool{
	"1_1":                                 true,
	"2_1":                                 true,
	"4_1":                                 true,
	"8_1":                                 true,
	"1_1-6_1":                             true,
	"2_1-3_1":                             true,
	"4_1-6_1":                             true,
	"1_1-2_1-8_1":                         true,
	"2_1-5_1-6_1":                         true,
	"1_1-2_1-5_1":                         true,
	"0_1-1_1-2_1-4_1":                     true,
	"0_1-2_1-4_1-8_1":                     true,
	"0_1-4_1-6_1-9_1":                     true,
	"1_1-2_1-8_1-9_1":                     true,
	"1_1-3_1-4_1-6_1-8_1":                 true,
	"2_1-3_1-6_1-7_1-8_1":                 true,
	"3_1-5_2-6_2":                         true,
	"0_1-1_2-2_1-3_1-7_1":                 true,
	"1_1-2_2-4_2-6_1":                     true,
	"2_2-4_1-5_1-8_2":                     true,
	"0_1-1_1-4_1-5_1-6_1-7_1-8_1":         true,
	"0_1-1_1-2_2-5_1-7_1-9_1":             true,
	"0_1-1_1-3_1-4_3-9_1":                 true,
	"0_1-3_1-6_1-8_4":                     true,
	"1_2-2_1-6_2-7_3":                     true,
	"2_1-3_3-4_2-5_2":                     true,
	"0_1-1_1-4_1-6_2-7_1-8_2":             true,
	"1_2-2_2-3_1-4_1-7_2-8_1":             true,
	"2_1-3_1-4_2-5_2-6_2-8_1":             true,
	"0_1-1_1-2_1-3_1-5_1-6_1-7_1-8_1-9_1": true,
}

func reorderedPowerOf2(N int) bool {
	// 检查 n 的编码是否在　isValid 中
	return isVaild[encode(N)]
}

func encode(n int) string {
	tmp := [10]int{}
	for n > 0 {
		tmp[n%10]++
		n /= 10
	}

	ss := make([]string, 0, 10)
	for n, c := range tmp {
		if c > 0 {
			ss = append(ss, fmt.Sprintf("%d_%d", n, c))
		}
	}

	return strings.Join(ss, "-")
}
