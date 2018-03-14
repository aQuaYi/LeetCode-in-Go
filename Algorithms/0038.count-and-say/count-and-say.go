package problem0038

func countAndSay(n int) string {
	buf := []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)
}

func say(buf []byte) []byte {
	// res 长度不会超过 buf 的两倍，所以，可以事先指定容量，加快append的速度
	res := make([]byte, 0, len(buf)*2)

	i, j := 0, 1
	for i < len(buf) {
		// 利用 j ，找到下一个不同的元素
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}

		// res 中 res[i] 表示 res[i+1] 的个数，i 为0,2,4,6,...
		res = append(res, byte(j-i+'0'), buf[i])

		// 移动 i 到 j
		i = j
	}

	return res
}
