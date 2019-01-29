package problem0190

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for bits := 0; bits < 32; bits++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
