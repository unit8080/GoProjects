// 1539. Kth Missing Positive Number
// https://leetcode.com/problems/kth-missing-positive-number/description/

func findKthPositive(arr []int, k int) int {
    
    left := 0
    right := len(arr) -1

    // binary search
    for left <= right {
        pivot := left + (right -left)/2
        if arr[pivot] - pivot -1  < k {
            left = pivot + 1 
        } else {
            right = pivot -1
        }
    }
    // left = right + 1
    // At the end of the loop, left = right + 1, 
    // and the kth missing number is in-between 
    // arr[right] and arr[left]. The number of integers 
    // missing before arr[right] is arr[right] - right - 1. 
    // Hence, the number to return is 
    // arr[right] + k - (arr[right] - right - 1) = k + left.

    return left + k
}
