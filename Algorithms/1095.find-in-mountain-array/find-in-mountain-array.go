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
 * This is the MountainArray's API interface.
 * You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, ma *MountainArray) int {
	n := ma.length()

	peak := searchPeak(0, n-1, ma)

	index := searchLeft(0, peak, target, ma)
	if index != -1 {
		return index
	}
	return searchRight(peak, n-1, target, ma)
}

func searchPeak(lo, hi int, ma *MountainArray) int {
	for lo < hi {
		mid := (lo + hi - 1) / 2
		if ma.get(mid) < ma.get(mid+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func searchLeft(lo, hi, target int, ma *MountainArray) int {
	less := func(a, b int) bool {
		return a < b
	}
	return binarySearch(lo, hi, target, ma, less)
}

func searchRight(lo, hi, target int, ma *MountainArray) int {
	less := func(a, b int) bool {
		return a > b
	}
	return binarySearch(lo, hi, target, ma, less)
}

func binarySearch(lo, hi, target int, ma *MountainArray, less func(a, b int) bool) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		m := ma.get(mid)
		if m == target {
			return mid
		} else if less(m, target) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
