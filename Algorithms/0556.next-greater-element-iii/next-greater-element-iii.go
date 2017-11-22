package Problem0556

import (
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	ss := make([]string, 0, 10)

	lastTail := n % 10
	isAvaliable := false
	for n > 0 {
		tail := n % 10
		if tail < lastTail {
			isAvaliable = true
		}
		lastTail = tail

		ss = append(ss, strconv.Itoa(tail))

		n /= 10
	}

	if !isAvaliable {
		return -1
	}

	reverse(ss)

	beg := exchange(ss)

	sort.Strings(ss[beg:])

	s := combine(ss)

	res, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return -1
	}
	return int(res)
}

func reverse(ss []string) {
	i, j := 0, len(ss)-1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}

func exchange(ss []string) int {
	var i, j int
	for j = len(ss) - 1; 1 <= j; j-- {
		for i = j - 1; 0 <= i; i-- {
			if ss[i] < ss[j] {
				ss[i], ss[j] = ss[j], ss[i]
				return i + 1
			}
		}
	}
	return -1
}

func combine(ss []string) string {
	str := ""
	for i := range ss {
		str += ss[i]
	}
	return str
}
