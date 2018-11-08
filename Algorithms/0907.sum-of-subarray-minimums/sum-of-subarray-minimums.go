package problem0907

const modulo = 1e9 + 7

// 求解 A 的所有子数组的最小值之和
func sumSubarrayMins(A []int) int {
	minOfA := 1 // 题目规定了 A[i]>=1
	A = append(A, minOfA-1)

	size := len(A)
	res := 0

	s := new(stack)
	s.push(0)

	for i := 1; i < size; i++ {
		for s.len() > 0 && A[s.top()] >= A[i] {
			j := s.pop()
			res += (j - s.top()) * (i - j) * A[j]
			// (j-s.top()) * (i-j) 是所有以 A[j] 为最小值的子数组的个数
			// j-s.top() 表示这些子数组可以有多少个左端点
			// i-j 表示这些子数组可以做多少个右端点
		}
		s.push(i)
	}

	return res % modulo
}

// stack 用于存放 A 中元素的 index
// 并且越往上的 index， A[index] 越大
type stack []int

func (s *stack) top() int {
	size := len(*s)
	if size == 0 {
		// 返回 -1 是为了使得 "res += ..." 行的逻辑统一
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
