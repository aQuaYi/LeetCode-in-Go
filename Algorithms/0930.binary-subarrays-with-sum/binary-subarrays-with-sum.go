package problem0930

func numSubarraysWithSum(A []int, S int) int {
	if S == 0 {
		return zero(A)
	}
	return nonZero(A, S)
}

func zero(A []int) int {
	res, count := 0, 0
	for _, n := range A {
		if n == 0 {
			count++
			res += count
		} else {
			count = 0
		}
	}
	return res
}

func nonZero(A []int, S int) int {
	size := len(A)

	end, sum := 0, 0
	for end < size && sum < S {
		sum += A[end]
		end++
	}

	if sum != S {
		// A 中所有的 1 加起来都比 S 小
		return 0
	}

	begin, res := 0, 0
	for begin < size && end <= size {
		left := 1
		for begin < size && A[begin] == 0 {
			begin++
			left++
		}

		right := 1
		for end < size && A[end] == 0 {
			end++
			right++
		}

		res += left * right

		// 此时， A[begin] == A[end] == 1
		begin++ // 相当于 sum--
		end++   // 相当于 sum++
		// 所以， sum 保持不变
	}

	return res
}
