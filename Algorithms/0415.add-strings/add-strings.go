package problem0415

func addStrings(s1, s2 string) string {
	// 确保 n1 <= n2
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	n1, n2 := len(s1), len(s2)
	a1, a2 := []byte(s1), []byte(s2)

	carry := byte(0)

	// buf 保存 []byte 格式的答案
	buf := make([]byte, n2+1)
	buf[0] = '1'

	i := 1
	for i <= n2 {
		// a1 和 a2 相加
		if i <= n1 {
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry

		// 处理进位问题
		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])
}
