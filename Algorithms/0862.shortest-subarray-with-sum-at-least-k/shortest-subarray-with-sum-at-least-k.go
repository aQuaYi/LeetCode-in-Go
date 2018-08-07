package problem0862

const intMax = 1<<63 - 1

func shortestSubarray(a []int, K int) int {
	size := len(a)

	sums := make([]int, size+1)
	s := 0
	for i, n := range a {
		if n == K {
			return 1
		}
		s += n
		sums[i+1] = s
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j <= size; j++ {
			if sums[j]-sums[i] == K {
				return j - i
			}
		}
	}

	return -1
}
