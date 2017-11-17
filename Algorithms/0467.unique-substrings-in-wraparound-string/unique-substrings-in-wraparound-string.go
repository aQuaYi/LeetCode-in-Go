package Problem0467

func findSubstringInWraproundString(p string) int {
	m := make(map[string]bool, len(p))

	isContinuous := func(i, j int) bool {
		a := int(p[i] - 'a')
		b := int(p[j] - 'a')
		return (a+j-i)%26 == b
	}

	var first, last int

	for last < len(p) {
		if !isContinuous(first, last) {
			first = last
			continue
		}

		for i := last; first <= i; i-- {
			m[p[i:last+1]] = true
		}

		last++
	}

	return len(m)
}
