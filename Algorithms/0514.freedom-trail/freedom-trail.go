package problem0514

func findRotateSteps(ring string, key string) int {
	// m 中记录了每个字符所有出现过的位置
	m := make(map[rune][]int, 26)
	for i, r := range ring {
		m[r] = append(m[r], i)
	}

	// ss 记录了到达同一个字符的各个位置的最佳状态
	//           	    0123456789
	// 例如， ring =	"acbcccccbd",
	// key = "ab"  时，0 → 2 是最优解
	// key = "abd" 时，0 → 8 → 9 才是最优解
	// 所以，需要记录达到同一个字符的各个位置的状态
	ss := make([]status, 1, len(ring))
	ss[0] = status{
		index: 0,
		steps: 0,
	}

	size := len(ring)

	for _, r := range key {
		temp := make([]status, len(m[r]))
		for i, d := range m[r] {
			// 寻找到达 r 字符的 d 位置的 steps 最小的解
			s := fromTo(ss[0], d, size)
			for i := 1; i < len(ss); i++ {
				// s.d 都是一样的，更新 s.steps即可
				s.steps = min(s.steps, fromTo(ss[i], d, size).steps)
			}
			temp[i] = s
		}
		ss = temp
	}

	res := ss[0].steps
	for i := 1; i < len(ss); i++ {
		res = min(res, ss[i].steps)
	}

	return res
}

// status 记录了当前的 12:00 方向对应的 ring 中字符的 index
// 和
// 已经操作的 steps
type status struct {
	index, steps int
}

// 从 s 状态到转动到 ring 的 d 位置，所生成的新状态
func fromTo(s status, d, size int) status {
	a := min(s.index, d)
	b := max(s.index, d)

	return status{
		index: d,
		steps: s.steps + min(b-a, size-(b-a)) + 1, // +1 表示按下
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
