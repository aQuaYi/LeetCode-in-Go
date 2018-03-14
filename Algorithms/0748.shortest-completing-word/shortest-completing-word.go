package problem0748

import (
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	keys := getKeys(licensePlate)
	minLen := 1<<63 - 1
	res := ""

	for _, w := range words {
		if len(w) >= minLen {
			continue
		}

		isCompleting := true
		for k, c := range keys {
			if strings.Count(w, k) < c {
				isCompleting = false
				break
			}
		}

		if isCompleting {
			res = w
			minLen = len(w)
		}
	}

	return res
}

func getKeys(licensePlate string) map[string]int {
	licensePlate = strings.ToLower(licensePlate)
	bs := []byte(licensePlate)

	res := make(map[string]int, len(licensePlate))
	for _, b := range bs {
		if 'a' <= b && b <= 'z' {
			res[string(b)]++
		}
	}

	return res
}
