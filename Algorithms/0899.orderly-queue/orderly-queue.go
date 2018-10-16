package problem0899

func orderlyQueue(s string, k int) string {
	seen := make(map[string]bool, len(s))
	min := s
	helper(s, k, &min, seen)
	return min
}

func helper(s string, k int, min *string, seen map[string]bool) {
	if seen[s] {
		return
	}

	seen[s] = true
	if *min > s {
		*min = s
	}

	for i := 0; i < k; i++ {
		helper(move(s, i), k, min, seen)
	}
}

func move(s string, i int) string {
	bytes := []byte(s)
	b := bytes[i]
	copy(bytes[i:], bytes[i+1:])
	bytes[len(bytes)-1] = b
	return string(bytes)
}
