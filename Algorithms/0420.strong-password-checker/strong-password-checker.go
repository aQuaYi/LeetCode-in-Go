package problem0420

func strongPasswordChecker(s string) int {
	// 统计缺少的字符类型
	l, u, d, missingType := 1, 1, 1, 3
	for i := range s {
		if 0 < l && 'a' <= s[i] && s[i] <= 'z' {
			l--
			missingType--
		}
		if 0 < u && 'A' <= s[i] && s[i] <= 'Z' {
			u--
			missingType--
		}
		if 0 < d && '0' <= s[i] && s[i] <= '9' {
			d--
			missingType--
		}

		if missingType == 0 {
			break
		}
	}

	var replace, ones, twos int

	for p := 0; p+2 < len(s); p++ {
		if s[p] != s[p+1] || s[p+1] != s[p+2] {
			continue
		}

		// 出现连续重复的情况
		// 统计连续重复的长度
		repeatingLen := 2
		for p+2 < len(s) && s[p] == s[p+2] {
			repeatingLen++
			p++
		}

		// 对于重复的字符串
		// 每3个字符，都需要替换为另外的字符
		replace += repeatingLen / 3
		if repeatingLen%3 == 0 {
			// 但是，如果字符串的长度超过了 20 的话
			// 对于 此种情况
			// 也可以通过 删除 1 个重复的字符，来减少替换的次数
			ones++
		} else if repeatingLen%3 == 1 {
			// 但是，如果字符串的长度超过了 20 的话
			// 对于 此种情况
			// 也可以通过 删除 2 个连续的重复的字符，来减少替换的次数
			twos++
		}
	}

	// 长度不够
	// 使用缺失的类型来凑
	if len(s) < 6 {
		return max(missingType, 6-len(s))
	}

	// 长度足够
	// 替换的时候，优先替换为缺失的类型
	if len(s) <= 20 {
		return max(missingType, replace)
	}

	// 长度太长了
	//
	// 需要删除的数量
	delete := len(s) - 20

	// 删除工作一定是要完成的，要不然长度要求就不能被满足了
	// 但是，删除那些连续重复的字符，可以减少 replace 的次数
	// 所以
	// 删除的时候，优先删除那些可以减少替换的连续重复字符

	// 通过删除 1 个重复的字符，可以减少 1 个 replace 的个数
	replace -= min(delete, ones)
	// 通过删除 2 个重复的字符，可以减少 1 个 replace 的个数
	replace -= min(max(delete-ones, 0), twos*2) / 2
	// 通过删除 3 个连续重复的字符，可以减少 1 个 replace 的个数
	replace -= max(delete-ones-2*twos, 0) / 3

	return delete + max(missingType, replace)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
