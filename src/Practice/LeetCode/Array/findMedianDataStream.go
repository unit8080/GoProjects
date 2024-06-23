// 295. Find Median from Data Stream
// https://leetcode.com/problems/find-median-from-data-stream/

type MaxHeap []int
type MinHeap []int

type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MinHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	index := len(*h) - 1
	old := *h
	x := old[index]
	*h = old[:index]
	return x
}
func (h *MinHeap) Pop() interface{} {
	index := len(*h) - 1
	old := *h
	x := old[index]
	*h = old[:index]
	return x
}
func (h MaxHeap) Top() int {
	return h[0]
}
func (h MinHeap) Top() int {
	return h[0]
}
func (this *MedianFinder) AddNum(num int) {
	if len(this.left)+len(this.right) == 0 {
		heap.Push(&this.left, num)
		return
	}
	for {
		if len(this.left) < len(this.right) {
			if num <= this.right.Top() {
				heap.Push(&this.left, num)
				return
			} else {
				v := heap.Pop(&this.right)
				heap.Push(&this.left, v)
			}
		} else {
			if num >= this.left.Top() {
				heap.Push(&this.right, num)
				return
			} else {
				v := heap.Pop(&this.left)
				heap.Push(&this.right, v)
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.left) == len(this.right) {
		return float64(this.left.Top()+this.right.Top()) / 2.0
	} else if len(this.left) > len(this.right) {
		return float64(this.left.Top())
	} else {
		return float64(this.right.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
