package problem0859

func buddyStrings(A, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		return hasDouble(A)
	}

	size := len(A)

	ca, cb := byte(0), byte(0)

	return false
}

func hasDouble(s string) bool {
	seen := [26]bool{}
	for i := range s {
		b := s[i] - 'a'
		if seen[b] {
			return true
		}
		seen[b] = true
	}
	return false
}
