package problem0753

import (
	"math"
	"strconv"
	"strings"
)

func crackSafe(n, k int) string {
	keys := makeKeys(n-1, k)
	r := make(map[string][]string, len(keys))
	for i := range keys {
		r[keys[i]] = makeNums(k)
	}

	res := strings.Repeat("0", n-1)

	count := int(math.Pow(float64(k), float64(n)))

	for k := 0; k < count; k++ {
		key := res[len(res)-(n-1):]
		next := r[key][0]
		r[key] = r[key][1:]
		res += next
	}

	return res
}

func makeKeys(keySize, k int) []string {
	nums := makeNums(k)
	res := []string{""}
	for s := 0; s < keySize; s++ {
		temp := make([]string, 0, len(res)*k)
		for i := range res {
			for j := range nums {
				temp = append(temp, res[i]+nums[j])
			}
		}
		res = temp
	}
	return res
}

func makeNums(k int) []string {
	res := make([]string, k)
	for i := range res {
		res[i] = strconv.Itoa(k - i - 1)
	}
	return res
}
