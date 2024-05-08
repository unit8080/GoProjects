// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/

/*
type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    maxHeap := &MaxHeap{}
    *maxHeap = nums
    heap.Init(maxHeap)

    i := 0
    last := math.MaxInt
    num := 0
    for i < k && len(*maxHeap) > 0 {
        num = (heap.Pop(maxHeap)).(int)
        if num != last {
            i += 1
        }
        // last = num
    }
    return num
}
*/


/*
func quickSortRecPivotFirstElement(arr []int , startIdx int, endIdx int) {
    // base case
    if startIdx >= endIdx {
        return
    }
    pivotIdx := startIdx
    leftIdx := startIdx+1
    rightIdx := endIdx
    
    for leftIdx <= rightIdx {
        if arr[leftIdx] > arr[pivotIdx] && 
        arr[pivotIdx] > arr[rightIdx] {
            swap(leftIdx, rightIdx, arr)
        }
        if arr[leftIdx] <= arr[pivotIdx] {
            leftIdx += 1
        }
        if arr[rightIdx] >= arr[pivotIdx] {
            rightIdx -= 1
        }
    }
    swap(pivotIdx, rightIdx, arr)
    quickSortRecPivotFirstElement(arr, startIdx, rightIdx - 1)
    quickSortRecPivotFirstElement(arr, rightIdx + 1, endIdx)
}
func quickSortPivotFirstElement(array []int) [] int {
    quickSortRecPivotFirstElement(array, 0, len(array) -1)
    return array
}
func findKthLargest(nums []int, k int) int {
    quickSortPivotFirstElement(nums)
    return nums[len(nums)-k]
}
*/

/*
func swap (i, j int, arr []int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, startIdx, endIdx int) int {
    pivot := arr[endIdx]
    partitionIdx := startIdx
    for i := startIdx; i < endIdx; i++ {
        if arr[i] <= pivot {
            swap(i, partitionIdx, arr)
            partitionIdx += 1
        }
    }
    swap(partitionIdx, endIdx, arr)
    return partitionIdx
}
var ans int
var result *int = &ans
func quickSort(arr []int , startIdx int, endIdx int, k int) {
    // base case
    if startIdx < endIdx {
        pIndex := partition(arr, startIdx, endIdx)
        if pIndex > k {
            quickSort(arr, startIdx, pIndex-1, k)
        } else if pIndex < k {
            quickSort(arr, pIndex+1, endIdx, k)
        } else {
            *result = arr[pIndex]
            return
        }
    } else {
        *result = arr[endIdx]
    }
}

func findKthLargest(nums []int, k int) int {
    quickSort(nums, 0, len(nums) -1,len(nums)-k)
    return *result
}
*/

/*
 - maxHeap with K-elements
 */

/*
 type MinHeap []int

 func (h MinHeap) Len() int {
     return len(h)
 }
 func (h MinHeap) Less(i, j int) bool {
     return h[i] < h[j]
 }
 func (h MinHeap) Swap(i, j int) {
     h[i], h[j] = h[j], h[i]
 }
 func (h *MinHeap) Push(x interface{}) {
     *h = append(*h, x.(int))
 }
 func (h *MinHeap) Pop() interface{} {
     old := *h
     n := len(old)
     x := old[n -1]
     *h = old[:n -1]
     return x
 }

func findKthLargest(nums []int, k int) int {
    minHeap := &MinHeap{}
    // *minHeap = nums
    // heap.Init(minHeap)
    // n := len(nums)
    i := 0
    for ; i < k; i += 1 {
        heap.Push(minHeap, nums[i])
    }
    // fmt.Println("minHeap:", *minHeap)
    
    for ; i < len(nums); i += 1 {
        old := *minHeap
        x := old[0]
        if nums[i] >= x {
             heap.Pop(minHeap)
            // fmt.Println("Popped:", pop)
            heap.Push(minHeap, nums[i])
            // fmt.Println("minHeap:", *minHeap)
        }
    }
    // fmt.Println("minHeap:", *minHeap)
    old := *minHeap
    x := old[0]
    return x
}
*/

func findKthLargest(nums []int, k int) int {
    
    partition := func(l, r int) int {
        for l < r {
            if nums[l+1] <= nums[l] {
                nums[l+1], nums[l] = nums[l], nums[l+1]
                l++
            } else if nums[r] >= nums[l] {
                r--
            } else {
                nums[r], nums[l+1] = nums[l+1], nums[r]
            }
        }
        return l
    }
    var qs func(l, r, n int) int 
    qs = func(l, r, n int) int {

        p := partition(l, r)
        if n-k < p {
            return qs(0, p-1, n)
        } else if n-k > p {
            return qs(p+1, r, n)
        } else {
            return p
        }
    }
    kth := qs(0, len(nums)-1, len(nums))
    return nums[kth]
}

