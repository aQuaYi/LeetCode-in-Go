package problem1146

// SnapshotArray can snap a array
type SnapshotArray struct {
	size      int
	snapCount int
	rec       map[int]map[int]int
}

// Constructor resturn s a SnapshotArray
func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		size: length,
	}
}

// Set val in index
func (sa *SnapshotArray) Set(index int, val int) {
	m, ok := sa.rec[sa.snapCount]
	if !ok {
		m = make(map[int]int, sa.size)
		sa.rec[sa.snapCount] = m
	}
	m[index] = val
}

// Snap make snapshot
func (sa *SnapshotArray) Snap() int {
	sa.snapCount++
	return sa.snapCount - 1
}

// Get returns val in the snap
func (sa *SnapshotArray) Get(index int, snapID int) int {
	return sa.rec[snapID][index]
}
