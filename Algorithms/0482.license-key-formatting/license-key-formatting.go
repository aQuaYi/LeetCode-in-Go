package Problem0482

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	countDashs := strings.Count(s, "-")
	countAN := len(s) - countDashs

	s = strings.Replace(s, "-", "", -1)
	s = strings.ToUpper(s)

	remain := countAN % k
	if remain == 0 {
		remain = k
	}

	res := make([]byte, (countAN+k-1)/k-1+countAN)

	i, j := len(res)-1, 1
	n := len(s)
	count := 0
	for {
		res[i] = s[n-j]
		j++
		count++
		i--

		if 0 <= i && count%k == 0 {
			res[i] = '-'
			i--
		}

		if i < 0 {
			break
		}
	}

	return string(res)
}
