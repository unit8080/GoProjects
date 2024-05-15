// 31. Next Permutation
// https://leetcode.com/problems/next-permutation/

func nextPermutation(nums []int)  {
    
    n := len(nums)
    pivot := n - 1

    for pivot >= 1 {
        if nums[pivot-1] >= nums[pivot] {
            pivot -= 1
        } else {
            break
        }
    }
    // fmt.Println("pivot:", pivot)
    if pivot > 0 {
        i := len(nums) -1
        for i >= 0  {
            if nums[i] <= nums[pivot -1] {
                i -= 1
            } else {
                break
            }
        }
        // swap here
        nums[pivot-1], nums[i] = nums[i], nums[pivot-1]
    }

    var reverse func(start, end int) 
    reverse = func(start, end int) {
        for start < end {
            nums[start], nums[end] = nums[end], nums[start]
            start++
            end--
        }
    }
    reverse(pivot, len(nums) -1) 
}

