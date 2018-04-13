package problem0793

func preimageSizeFZF(K int) int {
	for f := 305175781; f >= 1; f = (f - 1) / 5 {
		if K == 5*f {
			return 0
		}
		K %= f
	}

	return 5
}

// https://leetcode.com/problems/preimage-size-of-factorial-zeroes-function/discuss/117821/Four-binary-search-solutions-based-on-different-ideas
