// 705. Design HashSet
// https://leetcode.com/problems/design-hashset/

type MyHashSet struct {
    set map[interface{}]struct{}
}


func Constructor() MyHashSet {
    return MyHashSet {
        set: make(map[interface{}]struct{}),
    }
}


func (this *MyHashSet) Add(key int)  {
    set := this.set
    set[key] = struct{}{}
}


func (this *MyHashSet) Remove(key int)  {
    set := this.set
    delete(set, key)
}


func (this *MyHashSet) Contains(key int) bool {
    set := this.set
    if _, exists := set[key]; exists {
        return true
    } else {
        return false
    }
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
