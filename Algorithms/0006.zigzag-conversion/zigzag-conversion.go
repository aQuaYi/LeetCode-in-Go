package problem0006

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距
	p := numRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间的行
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}
