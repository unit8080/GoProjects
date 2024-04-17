/*
 * Algorithm:
 *  - API Put/Get keys needs O(1) processing in average case 
 *  - Requiremt to implement eviction policy
 *      - eviction policy to pick least recently used key,value pair
 *        when cache is upto capacity
 *      - so  we need to maintain list with most recently used at top and least
 *        recently used at the bottom, 
 *      - eviction should remove key from the tail
 *   - For O(1) add/delete, it needs key-value map and 
 *   - doubly linked list to maintain some sort of ordered keys whenever,
 *   - inserted or accessed
 *        
 */

type listNode struct {
    prev *listNode
    next *listNode
    key int
    value int
}

type LRUCache struct {
    capacity int
    head *listNode
    tail *listNode
    lruMap map[int] *listNode
    cacheMutex  sync.Mutex
}

func Constructor(capacity int) LRUCache {
    cache := new(LRUCache)
    cache.capacity = capacity
    cache.head = nil
    cache.tail = nil
    cache.lruMap = make(map[int] *listNode)
    return *cache
}

func (this *LRUCache) Get(key int) int {
    this.cacheMutex.Lock()
    defer this.cacheMutex.Unlock()
    
    node, exists := this.lruMap[key]
    if !exists {
        return -1
    }

    if this.head != node {
        this.unlinkNode(node)
        this.addToHead(node)
    }
    
    return node.value
}

func (this *LRUCache) Put(key int, value int)  {
    this.cacheMutex.Lock()
    defer this.cacheMutex.Unlock()

    node, exists := this.lruMap[key]
    if exists {
        this.removeItem(node)
    } 

    if len(this.lruMap) >= this.capacity {
        this.evictLRUItem()
    }
    newNode := new(listNode)
    newNode.key = key
    newNode.value = value
    this.addToHead(newNode)
    this.lruMap[key] = newNode
}

func (this *LRUCache) addToHead(node *listNode) {
    if node == nil {
        return
    }
    if this.head != nil {
        node.next = this.head
        this.head.prev = node
    }
    if this.tail == nil {
        this.tail = node
    }
    this.head = node
}

func (this *LRUCache) unlinkNode(node *listNode) {
    if node == nil { 
        return
    }
    prevNode := node.prev
    nextNode := node.next

    if prevNode != nil {
        prevNode.next = nextNode
    }
    if nextNode != nil {
        nextNode.prev = prevNode
    }
    if this.head == node {
        this.head = nextNode
    }
    if this.tail == node {
        this.tail = prevNode
    }
    node.prev = nil
    node.next = nil
}

func (this *LRUCache) removeItem(node *listNode) {
    this.unlinkNode(node)
    delete(this.lruMap, node.key)
}

func (this *LRUCache) evictLRUItem() {
    if this.tail != nil {
        node := this.tail
        this.removeItem(node)
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
