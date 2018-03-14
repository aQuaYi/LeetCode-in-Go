package problem0152

func maxProduct(a []int) int {
	cur, neg, max := 1, 1, a[0]

	for i := 0; i < len(a); i++ {

		switch {
		case a[i] > 0:
			cur, neg = a[i]*cur, a[i]*neg
		case a[i] < 0:
			cur, neg = a[i]*neg, a[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}
