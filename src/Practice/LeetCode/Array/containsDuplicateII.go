// 219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/

func containsNearbyDuplicate(nums []int, k int) bool {
    
    n := len(nums)

    set := map[int]bool{}
    i := 0
    j := 0
    for j < n {

        if (j - i) > k {
            delete(set, nums[i])
            i++
        }
        val, _ := set[nums[j]]
        if val {
            return true
        }
        set[nums[j]] = true
        
        j++
    }
    return false
}
