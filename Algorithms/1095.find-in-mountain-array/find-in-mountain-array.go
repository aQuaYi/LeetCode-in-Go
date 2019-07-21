package problem1095

// MountainArray is required
type MountainArray []int

func (m *MountainArray) get(index int) int {
	return (*m)[index]
}
func (m *MountainArray) length() int {
	return len(*m)
}

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, ma *MountainArray) int {
	n := ma.length()
	lo, hi := 1, n-2
	for lo < hi {
		mid := (lo + hi - 1) / 2
		if ma.get(mid-1) < ma.get(mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	high := lo - 1
	mh := ma.get(high)

	if mh < target {
		return -1
	}

	if mh == target {
		return high
	}

	m0 := ma.get(0)
	if m0 == target {
		return 0
	}

	if m0 < target {
		lo, hi := 1, high-1
		for lo <= hi {
			mid := (lo + hi) / 2
			mm := ma.get(mid)
			if mm == target {
				return mid
			} else if mm < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return -1
	}

	me := ma.get(n - 1)
	if me == target {
		return n - 1
	}

	lo, hi = high+1, n-2
	for lo <= hi {
		mid := (lo + hi) / 2
		mm := ma.get(mid)
		if mm == target {
			return mid
		} else if mm > target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1

}
