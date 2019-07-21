package problem1095

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
func findInMountainArray(target int, mountainArr *MountainArray) int {

	return 0
}

// MountainArray is required
type MountainArray []int

func (m *MountainArray) get(index int) int {
	return (*m)[index]
}
func (m *MountainArray) length() int {
	return len(*m)
}
