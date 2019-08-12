package problem1146

// SnapshotArray can snap a array
type SnapshotArray struct {
	snaps   []map[int]int
	current map[int]int
	id      int
}

// Constructor resturn s a SnapshotArray
func Constructor(length int) SnapshotArray {
	s := make([]map[int]int, 0, 128)
	c := make(map[int]int, 32)
	return SnapshotArray{
		snaps:   s,
		current: c,
		id:      0,
	}
}

// Set val in index
func (sa *SnapshotArray) Set(index int, val int) {
	sa.current[index] = val
}

// Snap make snapshot
func (sa *SnapshotArray) Snap() int {
	sa.snaps = append(sa.snaps, sa.current)
	sa.current = make(map[int]int)
	sa.id++
	return sa.id - 1
}

// Get returns val in the snap
func (sa *SnapshotArray) Get(index int, snap int) int {
	for id := snap; id >= 0; id-- {
		history := sa.snaps[id]
		if t, ok := history[index]; ok {
			return t
		}
	}
	return 0 //default value
}
