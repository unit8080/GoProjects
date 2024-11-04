// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/?envType=company&envId=nvidia&favoriteSlug=nvidia-three-months


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

type RandomizedSet struct {
    set map[int]int 
    arr []int 
}


func Constructor() RandomizedSet {
    randomSet := new(RandomizedSet)
    randomSet.set = map[int]int{}
    randomSet.arr = []int{}
    // rand.Seed(time.Now().UnixNano())
    return *randomSet
}


func (this *RandomizedSet) Insert(val int) bool {

    _, exists := this.set[val]
    if exists {
        return false
    }
    index := len(this.arr)
    this.arr = append(this.arr, val)
    this.set[val] = index
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    index, exists := this.set[val]
    if !exists {
        return false
    }
    lastIndex := len(this.arr) -1
    lastVal := this.arr[lastIndex]
    delete(this.set, val)
    this.arr = this.arr[:lastIndex]
    if lastIndex == index {
        return true
    }
    this.set[lastVal] = index
    this.arr[index] = lastVal
    return true
}


func (this *RandomizedSet) GetRandom() int {
    n := len(this.set)
    randomNumber := rand.Intn(n)
    return this.arr[randomNumber]
}

