package Problem0344

func reverseString(s string) string {
	bytes := []byte(s)
	size := len(s)
	res := make([]byte, size)

	for i := range bytes {
		res[i] = bytes[size-i-1]
	}

	return string(res)
}
