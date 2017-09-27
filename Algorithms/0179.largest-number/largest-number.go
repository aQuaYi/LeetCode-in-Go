package Problem0179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	size := len(nums)
	b := make(bytes, size)
	resSize := 0
	for i := range b {
		b[i] = []byte(strconv.Itoa(nums[i]))
		resSize += len(b[i])
	}

	sort.Sort(b)

	res := make([]byte, 0, resSize)
	for i := size - 1; i >= 0; i-- {
		res = append(res, b[i]...)
	}

	return string(res)
}

type bytes [][]byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	bi, bj := b[i], b[j]
	if len(bi) != len(bj) {
		bi, bj = toSameLen(bi, bj)
	}
	size := len(bi)
	for k := 0; k < size; k++ {
		if bi[k] < bj[k] {
			return true
		}
	}

	return false
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func toSameLen(a, b []byte) ([]byte, []byte) {
	if len(a) > len(b) {
		tail := make([]byte, len(a)-len(b))
		for i := range tail {
			tail[i] = b[len(b)-1]
		}

		return a, append(b, tail...)
	}

	rb, ra := toSameLen(b, a)
	return ra, rb
}
