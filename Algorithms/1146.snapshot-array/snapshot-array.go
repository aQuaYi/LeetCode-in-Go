package problem1146

// SnapshotArray can snap a array
type SnapshotArray struct {
	count int
	rec   []map[int]int
}

// Constructor resturn s a SnapshotArray
func Constructor(length int) SnapshotArray {
	rec := make([]map[int]int, length)
	for i := range rec {
		m := make(map[int]int, 32)
		rec[i] = m
	}
	return SnapshotArray{
		rec: rec,
	}
}

// Set val in index
func (sa *SnapshotArray) Set(index int, val int) {
	sa.rec[index][sa.count] = val
}

// Snap make snapshot
func (sa *SnapshotArray) Snap() int {
	sa.count++
	return sa.count - 1
}

// Get returns val in the snap
func (sa *SnapshotArray) Get(index int, snap int) int {
	rec := sa.rec[index]
	res, ok := rec[snap]
	for !ok && snap > 0 {
		snap--
		res, ok = rec[snap]
	}
	if ok {
		return res
	}
	return 0
}
