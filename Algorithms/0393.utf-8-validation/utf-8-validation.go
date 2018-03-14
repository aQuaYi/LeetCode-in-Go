package problem0393

func validUtf8(data []int) bool {
	// cnt 是还需检查的 byte 的个数
	cnt := 0
	var d int
	for _, d = range data {
		if cnt == 0 {
			switch {
			case d>>3 == 30: //0b11110
				cnt = 3
			case d>>4 == 14: //0b1110
				cnt = 2
			case d>>5 == 6: //0b110
				cnt = 1
			case d>>7 > 0:
				// data[0] 和 data[len(data)-1] 都会到此处来检查
				return false
			}
		} else {
			// 非首尾的 byte 必须以 0b10 开头
			if d>>6 != 2 { //0b10
				return false
			}
			cnt--
		}
	}

	return 0 == cnt
}
