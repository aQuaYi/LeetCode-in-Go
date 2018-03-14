package problem0670

import (
	"strconv"
)

func maximumSwap(num int) int {
	bs := []byte(strconv.Itoa(num))

	// indexs 记录了每个数字最后出现的位置
	indexs := make(map[byte]int, len(bs))
	for i := range bs {
		indexs[bs[i]] = i
	}

	// 高位上的数字与低位上的更大的数字交换，才能使得 num 变大
	for i := 0; i < len(bs); i++ {
		for bj := byte('9'); bs[i] < bj; bj-- {
			j := indexs[bj]
			if j > i {
				bs[i], bs[j] = bs[j], bs[i]
				return convert(bs)
			}
		}
	}

	return convert(bs)
}

func convert(bs []byte) int {
	n, _ := strconv.Atoi(string(bs))
	return n
}
