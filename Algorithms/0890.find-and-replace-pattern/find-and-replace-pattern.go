package problem0890

func findAndReplacePattern(words []string, pattern string) []string {
	res := make([]string, 0, len(words))
	for _, w := range words {
		if isOK(w, pattern) {
			res = append(res, w)
		}
	}
	return res
}

func isOK(word, parttern string) bool {
	size := len(parttern)
	m2w := make(map[byte]byte, size)
	m2p := make(map[byte]byte, size)
	for i := 0; i < size; i++ {
		wi := word[i]
		pi := parttern[i]
		mwi, wok := m2w[pi]
		mpi, pok := m2p[wi]

		if !wok && !pok {
			m2w[pi] = wi
			m2p[wi] = pi
		} else if wi != mwi || pi != mpi {
			return false
		}
	}
	return true
}
