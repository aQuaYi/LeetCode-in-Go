package problem1146

// SnapshotArray can snap a array
type SnapshotArray struct {
	array []int
	snaps [][]int
}

// Constructor resturn s a SnapshotArray
func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		array: make([]int, length),
	}
}

// Set val in index
func (sa *SnapshotArray) Set(index int, val int) {
	sa.array[index] = val
}

// Snap make snapshot
func (sa *SnapshotArray) Snap() int {
	sa.snaps = append(sa.snaps, append(sa.array[:0:0], sa.array...))
	return len(sa.snaps) - 1
}

// Get returns val in the snap
func (sa *SnapshotArray) Get(index int, snapID int) int {
	return sa.snaps[snapID][index]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
