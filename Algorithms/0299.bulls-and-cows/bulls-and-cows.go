package Problem0299

import "fmt"

func getHint(secret string, guess string) string {
	var A, B int
	s := []byte(secret)
	g := []byte(guess)

	// 寻找 bulls
	for i := range s {
		if g[i] == s[i] {
			A++
			g[i] = 'x'
			s[i] = 'x'
		}
	}

	// 寻找 cows
	for i := range s {
		if s[i] == 'x' {
			continue
		}

		for j := range g {
			if g[j] == 'x' {
				continue
			}

			if s[i] == g[j] {
				B++
				g[j] = 'x'
				break
			}
		}
	}

	return fmt.Sprintf("%dA%dB", A, B)
}
