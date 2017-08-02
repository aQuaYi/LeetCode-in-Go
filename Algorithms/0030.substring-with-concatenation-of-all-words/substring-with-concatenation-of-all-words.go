package Problem0030

import "sort"

func findSubstring(s string, words []string) []int {
	lenWord := len(words[0])

	if lenWord == 0 {
		res := make([]int, len(s)+1)
		for i := range res {
			res[i] = i
		}
		return res
	}

	res := []int{}

	record := getRecord(s, words)
	numWords := len(words)

	for i, r := range record {
		for j := 0; j <= len(r)-numWords; j++ {

			temp := append([]int{}, r[j:j+numWords]...)
			if isFound(temp) {
				res = append(res, i+j*lenWord)
			}
		}
	}

	return res
}

func getRecord(s string, words []string) map[int][]int {
	lenWord := len(words[0])

	record := map[int][]int{}
	for i := 0; i < lenWord; i++ {
		addRecord(s[i:], words, i, &record)
	}

	return record
}

func addRecord(s string, words []string, index int, record *map[int][]int) {
	lenWord := len(words[0])
	raw := make([]int, len(s)/lenWord)

	for i := 0; i <= len(s)-lenWord; i += lenWord {
		raw[i/lenWord] = getIndex(words, s[i:i+lenWord])
	}

	i := 0
	for i < len(raw) {
		for i < len(raw) && raw[i] == -1 {
			i++
		}
		temp := []int{}
		localIndex := i
		for i < len(raw) && raw[i] != -1 {
			temp = append(temp, raw[i])
			i++
		}
		if len(temp) >= len(words) {
			(*record)[index+localIndex*lenWord] = temp
		}
	}
}

func isFound(nums []int) bool {
	sort.Ints(nums)
	for i, v := range nums {
		if v != i {
			return false
		}
	}

	return true
}

func getIndex(strs []string, str string) int {
	for i, s := range strs {
		if s == str {
			return i
		}
	}

	return -1
}
