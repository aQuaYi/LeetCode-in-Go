package problem0907

const modulo = 1e9 + 7

// 求解 A 的所有子数组的最小值之和
func sumSubarrayMins(A []int) int {
	minOfA := 1             // 题目规定了 A[i]>=1
	A = append(A, minOfA-1) // 加入低过 A 下限的值，是为了简化代码

	size := len(A)
	res := 0

	s := new(stack)
	s.push(0)

	for k := 1; k < size; k++ {
		for s.len() > 0 && A[s.top()] > A[k] {
			j := s.pop()
			i := s.top()
			res += (j - i) * (k - j) * A[j]
			// 由题意可知，
			// 为了找到全部以 A[j] 为最小值的子数组，
			// 需要找到这些子数组左右端点的个数
			// anyoneOf(A[i+1:j]) > A[j]，但是 A[i] < A[j]
			// 	   所以，A[i+1:j+1] 中的数，都可以做左端点，共有 j-i 个
			// anyoneOf(A[j+1:k]) > A[j]，但是 A[j] > A[k]
			//     所以，A[j:k] 中的数，都可以做右端点，共有 k-j 个
			// 两边的端点可以任意组合，所以是乘法关系
			// (j-i) * (k-j) 是所有以 A[j] 为最小值的子数组的个数
		}
		s.push(k)
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
