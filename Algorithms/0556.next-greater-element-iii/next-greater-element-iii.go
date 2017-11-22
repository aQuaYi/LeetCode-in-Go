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
	maxStr := string('9' + 1)

	for i = len(ss) - 2; 0 <= i; i-- {
		n := ss[i]
		minGreater := maxStr
		minIndex := i
		for j = i + 1; j < len(ss); j++ {
			if n < ss[j] && ss[j] < minGreater {
				minGreater = ss[j]
				minIndex = j
			}
		}

		if i < minIndex {
			ss[i], ss[minIndex] = ss[minIndex], ss[i]
			return i + 1
		}
	}

	// 因为主函数前面的有 isAValiable 检查过了
	// 所以，总会出现交换的情况
	panic("NEVER BE HERE")
}

func combine(ss []string) string {
	str := ""
	for i := range ss {
		str += ss[i]
	}
	return str
}
