package problem0537

import (
	"fmt"
	"strconv"
	"strings"
)

// c1 = a + b*i
// c2 = c + d*i
// res = (a*c-b*d) + (a*d+b*c)*i
func complexNumberMultiply(c1 string, c2 string) string {
	a, b := analyze(c1)
	c, d := analyze(c2)

	return fmt.Sprintf("%d+%di", a*c-b*d, a*d+b*c)
}

// c == re + im*i
// 返回 re 和 im
func analyze(c string) (re, im int) {
	ss := strings.Split(c, "+")
	re, _ = strconv.Atoi(ss[0])
	im, _ = strconv.Atoi(ss[1][:len(ss[1])-1])
	return
}
