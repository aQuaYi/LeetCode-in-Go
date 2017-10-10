package Problem0299

import "strconv"

func getHint(secret string, guess string) string {
	var bulls, cows int
	var counts [10]int

	size := len(secret)
	for i := 0; i < size; i++ {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')

		if ns == ng {
			bulls++
		} else {
			if counts[ns] < 0 {
				cows++
			}
			counts[ns]++

			if counts[ng] > 0 {
				cows++
			}
			counts[ng]--
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
