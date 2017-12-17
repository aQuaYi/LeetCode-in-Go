package Problem0649

var result = map[byte]string{
	'R': "Radiant",
	'D': "Dire",
}

func predictPartyVictory(senate string) string {
	bs := make([]byte, 0, len(senate)*2)
	bs = append(bs, []byte(senate)...)
	return helper(bs)
}

func helper(bs []byte) string {
	b := bs[0]
	var i int
	for i = 1; i < len(bs); i++ {
		if bs[i] != b {
			break
		}
	}

	if i == len(bs) {
		return result[b]
	}

	copy(bs[i:], bs[i+1:])
	bs[len(bs)-1] = b
	return helper(bs[1:])
}
