package problem0869

// 对 [1, 1e9] 中所有的 2^n 进行编码
var isPowerOf2 = map[int]bool{
	0:           true, // 2^0
	256:         true, // 2^1
	65536:       true, // 2^2
	4294967296:  true, // 2^3
	16777216:    true, // 2^4
	4352:        true, // 2^5
	16842752:    true, // 2^6
	4294967552:  true, // 2^7
	17826048:    true, // 2^8
	1048848:     true, // 2^9
	65793:       true, // 2^10
	4295033089:  true, // 2^11
	68736319489: true, // 2^12
	73014444304: true, // 2^13
	4311814144:  true, // 2^14
	4580184320:  true, // 2^15
	35655680:    true, // 2^16
	268439825:   true, // 2^17
	16908816:    true, // 2^18
	8591049216:  true, // 2^19
	4581294081:  true, // 2^20
	68988961297: true, // 2^21
	68719677457: true, // 2^22
	17196650497: true, // 2^23
	838861072:   true, // 2^24
	2240768:     true, // 2^25
	8891990033:  true, // 2^26
	4831908368:  true, // 2^27
	4330754304:  true, // 2^28
	73300709649: true, // 2^29
}

func reorderedPowerOf2(N int) bool {
	return isPowerOf2[encode(N)]
}

// 由于 n 的范围是 [1,1e9]
// 每个数字出现的次数不会超过 10 次，可以用 4 bit 的二进制数记录
// go 的 int 类型是 64 bit，足够记录 10 个数字的出现次数
// 所以， res 的 [4*digit, 4*digit+4) 位代表的二进制数，表示 digit 在 n 中出现的次数
func encode(n int) int {
	res := 0
	for n > 1 {
		digit := uint(n % 10)
		res += 1 << (4 * digit)
		n /= 10
	}
	return res
}
