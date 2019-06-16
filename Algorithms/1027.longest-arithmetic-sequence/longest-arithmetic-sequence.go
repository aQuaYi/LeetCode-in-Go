package problem1027

func longestArithSeqLength(A []int) int {
	n := len(A)
	exist, same := make(map[int][]int, n), make(map[int]int, n)
	res := 0
	for i := 0; i < n; i++ {
		if _, ok := exist[A[i]]; ok {
			exist[A[i]] = append(exist[A[i]], i)
		} else {
			exist[A[i]] = []int{i}
		}
		same[A[i]]++
		if same[A[i]] > res {
			res = same[A[i]]
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := A[j] - A[i]
			if diff != 0 {
				p, index := 2, j
				for {
					if indexes, ok := exist[A[index]+diff]; ok {
						breakFlag := true
						for k := 0; k < len(indexes); k++ {
							if indexes[k] > index {
								p++
								index = indexes[k]
								breakFlag = false
								break
							}
						}

						if breakFlag {
							break
						}
					} else {
						break
					}
				}

				if p > res {
					res = p
				}
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
