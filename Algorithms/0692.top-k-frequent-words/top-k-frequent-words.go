package problem0692

import "sort"

func topKFrequent(words []string, k int) []string {
	count := make(map[string]int, len(words))
	for _, w := range words {
		count[w]++
	}

	fw := make(freWords, 0, len(count))
	for w, c := range count {
		fw = append(fw, &entry{
			word:      w,
			frequence: c,
		})
	}

	sort.Sort(fw)

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = fw[i].word
	}

	return res
}

// entry 是 priorityQueue 中的元素
type entry struct {
	word      string
	frequence int
}

// PQ implements heap.Interface and holds entries.
type freWords []*entry

func (f freWords) Len() int { return len(f) }

func (f freWords) Less(i, j int) bool {
	if f[i].frequence == f[j].frequence {
		return f[i].word < f[j].word
	}
	return f[i].frequence > f[j].frequence
}

func (f freWords) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
