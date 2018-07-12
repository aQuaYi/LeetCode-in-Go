package problem0859

func buddyStrings(A, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		return hasDouble(A)
	}

	size := len(A)
	i := 0
	countDown := 2
	ca, cb := byte(0), byte(0)
	for countDown > 0 && i < size {
		if A[i] != B[i] {
			ca += A[i]
			cb += B[i]
			countDown--
		}
		i++
	}

	return ca == cb && A[i:] == B[i:]
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
