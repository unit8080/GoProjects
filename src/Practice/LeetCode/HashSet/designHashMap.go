// 706. Design HashMap
// https://leetcode.com/problems/design-hashmap/

type MyHashMap struct {
    hash map[interface{}]interface{}
}


func Constructor() MyHashMap {
    return MyHashMap{
        hash: make(map[interface{}]interface{}),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    hash := this.hash
    hash[key] = value
}


func (this *MyHashMap) Get(key int) int {
    hash := this.hash
    if val, exists := hash[key]; exists {
        return val.(int)
    } else {
        return -1
    }
}


func (this *MyHashMap) Remove(key int)  {
    hash := this.hash
    delete(hash, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
