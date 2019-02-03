package problem0936

import (
	"container/list"
	"strings"
)

func movesToStamp(stamp string, target string) []int {
	r := strings.Repeat("*", len(stamp))
	l := list.New()
	for strings.Contains(target, stamp) {
		index := strings.Index(target, stamp)
		l.PushBack(index)
		target = strings.Replace(target, stamp, r, 1)
	}

	return nil
}
