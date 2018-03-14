package problem0318

func maxProduct(words []string) int {
	size := len(words)

	masks := make([]int, size)
	for i := 0; i < size; i++ {
		for _, b := range words[i] {
			// 利用 bite 位来描述 words[i]
			masks[i] |= (1 << uint32(b-'a'))
		}
	}

	var max int
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			// 当 words[i] 和 words[j] 拥有相同的字母时
			// 这个字母所对应的位进行 & 运算后，为 1
			if masks[i]&masks[j] != 0 {
				continue
			}
			temp := len(words[i]) * len(words[j])
			if max < temp {
				max = temp
			}
		}
	}

	return max
}
