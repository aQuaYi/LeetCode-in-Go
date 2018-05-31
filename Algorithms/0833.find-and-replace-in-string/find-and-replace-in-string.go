package problem0833

import (
	"bytes"
	"sort"
)

type entry struct {
	idx            int
	source, target string
}

func findReplaceString(Str string, indexes []int, sources []string, targets []string) string {
	es := make([]*entry, len(indexes))
	for i := range es {
		es[i] = &entry{
			idx:    indexes[i],
			source: sources[i],
			target: targets[i],
		}
	}

	// 按照 idx 升排序
	sort.Slice(es, func(i int, j int) bool {
		return es[i].idx < es[j].idx
	})

	size := len(Str)
	ss := make([]string, 0, len(indexes)*2+1)

	end := 0
	// 按照 es 分割 Str
	for _, e := range es {
		begin := end
		ss = append(ss, Str[begin:e.idx])
		end = min(size, e.idx+len(e.source))
		ss = append(ss, Str[e.idx:end])
	}
	// 尾部有剩余，也要加入 ss
	if end < size {
		ss = append(ss, Str[end:])
	}

	var buf bytes.Buffer

	for i, s := range ss {
		// ss 索引号为偶数的元素，不需要替换
		if i%2 == 0 {
			buf.WriteString(s)
			continue
		}

		// ss 索引号为奇数的元素，需要检查是否要替换
		if s == es[i/2].source {
			buf.WriteString(es[i/2].target)
		} else {
			buf.WriteString(s)
		}
	}

	return buf.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
