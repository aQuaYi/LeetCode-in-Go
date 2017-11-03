package Problem0402

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	res := make([]byte, len(num)-k)
	i, j := len(num)-1, len(res)-1
	for 0 < k && 0 < i && 0 <= j {
		if num[i-1] > num[i] {
			res[j] = num[i]
			i--
			j--
			continue
		}
		i--
		k--
	}
	if 0 < j {
		copy(res[:j], num[:i])
	}

	i = 0
	for i < len(res)-1 && res[i] == '0' {
		i++
	}
	return string(res[i:])
}
