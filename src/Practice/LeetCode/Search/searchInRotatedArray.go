// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
    n := len(nums)
    if n < 1 {
        return -1
    }
    
    pivot := findPivot(nums, 0, n -1) 

    idx := searchBinary(nums, 0, pivot -1, target)
    if idx == -1 {
        idx = searchBinary(nums, pivot, n -1, target)
    }
    return idx
}

func searchBinary(nums []int, l, r, target int) int {
    for l <= r {
        mid := l + (r -l)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            r = mid -1
        } else {
            l = mid +1
        }
    }
    return -1
}

func findPivot(nums []int, l, r int) int {
    for l < r {
        mid := l + (r -l)/2
        if nums[mid] > nums[r] {
            l = mid +1
        } else {
            r = mid
        }
    }
    return r
}


