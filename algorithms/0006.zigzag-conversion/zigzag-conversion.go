package Problem0006

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// 准备存放字符的容器
	temp := make([]bytes.Buffer, numRows)
	// 准备序列号生成函数
	index := indexFunc(numRows)

	// 按照序列号生成的方式，
	// 依次把字符存入指定的容器
	for _, c := range s {
		i := index()
		temp[i].WriteRune(c)
	}

	// 汇总各个容器的内容
	res := ""
	for i := range temp {
		res += temp[i].String()
	}

	return res
}

// indexFunc 返回一个函数index
// index()会产生所需的容器序号
// 当 num == 4 时
// indexS == []int{0,1,2,3,2,1}
func indexFunc(num int) func() int {
	len := num*2 - 2
	indexS := make([]int, len)

	for i := 0; i < len; i++ {
		if i < num {
			indexS[i] = i
		} else {
			indexS[i] = indexS[len-i]
		}
	}

	i := -1
	return func() int {
		i++
		return indexS[i%len]
	}
}
