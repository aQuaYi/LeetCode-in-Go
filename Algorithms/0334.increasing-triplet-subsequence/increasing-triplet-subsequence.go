package Problem0334

func increasingTriplet(nums []int) bool {
	var i, n int
	rec := [][]int{[]int{nums[0]}}
	for _, n = range nums {
		for i = range rec {
			if n > rec[i][0] {
				if n < rec[i][len(rec[i])-1] {
					rec[i][len(rec[i])-1] = n
				} else if n > rec[i][len(rec[i])-1] {
					if len(rec[i]) == 2 {
						return true
					}
					rec[i] = append(rec[i], n)
				}
			}
		}

		if n < rec[len(rec)-1][0] {
			temp := make([]int, 1, 2)
			temp[0] = n
			rec = append(rec, temp)
		}
	}

	return false
}
