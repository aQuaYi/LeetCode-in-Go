package Problem0423

import (
	"strings"
)

type digit struct {
	b       byte
	key     string
	letters []string
}

var digits = []digit{
	{'2', "w", []string{"t", "w", "o"}},
	{'0', "z", []string{"z", "e", "r", "o"}},
	{'8', "g", []string{"e", "i", "g", "h", "t"}},
	{'3', "h", []string{"t", "h", "r", "e", "e"}},
	{'6', "x", []string{"s", "i", "x"}},
	{'7', "s", []string{"s", "e", "v", "e", "n"}},
	{'5', "v", []string{"f", "i", "v", "e"}},
	{'4', "f", []string{"f", "o", "u", "r"}},
	{'1', "o", []string{"o", "n", "e"}},
}

func originalDigits(s string) string {
	size := 0
	count := [10]int{}

	countFor := func(d digit) {
		c := strings.Count(s, d.key)
		if c == 0 {
			return
		}
		count[d.b-'0'] = c
		size += c
		for _, l := range d.letters {
			s = strings.Replace(s, l, "", c)
		}
	}

	for _, d := range digits {
		countFor(d)
	}

	count[9] = len(s) / 4
	size += count[9]

	res := make([]byte, 0, size)
	for b := byte(0); b <= 9; b++ {
		for count[b] > 0 {
			res = append(res, b+'0')
			count[b]--
		}
	}
	return string(res)
}
