package Problem0013

import (
	"strings"
)

func romanToInt(s string) int {

	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}

	res := 0
	base := 1000
	for i := 3; i >= 0; i-- {
		for j := len(d[i]) - 1; j > 0; j-- {
			if strings.HasPrefix(s, d[i][j]) {
				s = strings.TrimPrefix(s, d[i][j])
				res += base * j
				break
			}
		}
		base /= 10
	}

	return res
}
