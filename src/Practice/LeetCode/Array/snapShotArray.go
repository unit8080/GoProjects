// 1146. Snapshot Array
// https://leetcode.com/problems/snapshot-array/

type SnapshotArray struct {
    snapId int
    history map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
    snap := new(SnapshotArray)
    snap.snapId = 0
    snap.history = make(map[int]map[int]int, length)
    return *snap
}

func (this *SnapshotArray) Set(index int, val int)  {
    if _, ok := this.history[index]; !ok {
        this.history[index] = make(map[int]int)
    }
    this.history[index][this.snapId] = val
}

func (this *SnapshotArray) Snap() int {
    this.snapId++
    return this.snapId -1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
    m := this.history[index]

    for snap_id >= 0 {
        if _, ok := m[snap_id]; !ok {
            snap_id--
        } else {
            break
        }
    }
    
    if this.snapId < 0 {
		return 0
	}
    return m[snap_id]
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
