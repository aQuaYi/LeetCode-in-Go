package Problem0402

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	res := make([]byte, len(num)-k)
	var i, j int
	for 0 < k && i+1 < len(num) && j < len(res) {
		if num[i] <= num[i+1] {
			res[j] = num[i]
			i++
			j++
			continue
		}
		i++
		k--
	}
	copy(res[j:], num[i:])

	i = 0
	for i < len(res)-1 && res[i] == '0' {
		i++
	}
	return string(res[i:])
}
