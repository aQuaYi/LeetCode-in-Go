package problem0443

import "strconv"

func compress(chars []byte) int {
	end, count := 0, 1

	for i, char := range chars {
		if i+1 < len(chars) && char == chars[i+1] {
			// 统计相同字符的个数
			count++
		} else { // 出现不同的字符
			// 在 end 处填写被压缩的字符
			chars[end] = char
			end++

			// 从 end 处填写被压缩字符的个数
			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[end] = byte(num)
					end++
				}
			}

			// 把 count 重置为 1
			count = 1
		}
	}

	return end
}
