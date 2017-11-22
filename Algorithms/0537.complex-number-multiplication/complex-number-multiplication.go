package Problem0537

import (
	"strconv"
	"strings"
)

// c1 = a + b*i
// c2 = c + d*i
// res = (ac-bd) + (ad+bc)*i
func complexNumberMultiply(c1 string, c2 string) string {

	return res
}

func get(c string) (re, im int) {
	ss := strings.Split(c, "+")
	re, _ = strconv.Atoi(ss[0])
	im, _ = strconv.Atoi(ss[1][:len(ss[1]-2)])
	return
}
