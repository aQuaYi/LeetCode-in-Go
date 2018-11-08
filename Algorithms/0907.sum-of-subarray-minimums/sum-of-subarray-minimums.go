package problem0907

const modulo = 1e9 + 7

// 求解 A 的所有子数组的最小值之和
func sumSubarrayMins(A []int) int {
	s := new(stack)
	s.push(0)

	// A = append(A, 0)
	size := len(A)
	res := 0

	for j := 1; j < size; j++ {
		if A[s.top()] < A[j] {
			s.push(j)
			continue
		}

		for s.len() > 0 && A[s.top()] >= A[j] {
			i := s.pop()
			res += (i - s.top()) * (j - i) * A[i]
		}

		s.push(j)
	}

	for s.len() > 0 {
		i := s.pop()
		res += (i - s.top()) * (size - i) * A[i]
	}

	return res % modulo
}

// stack 用于存放 A 中元素的 index
// 并且越往上的 index， A[index] 越大
type stack []int

func (s *stack) top() int {
	size := len(*s)
	if size == 0 {
		return -1
	}
	return (*s)[size-1]
}

func (s *stack) push(index int) {
	*s = append(*s, index)
}

func (s *stack) pop() (index int) {
	size := len(*s)
	*s, index = (*s)[:size-1], (*s)[size-1]
	return
}

func (s *stack) len() int {
	return len(*s)
}
