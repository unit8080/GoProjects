// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
    
    set := make(map[int] struct{})

    for i := 0; i < len(nums); i++ { 
        if _, exists := set[nums[i]]; exists {
            return true
        } else {
            set[nums[i]] = struct{}{}
        }
    }
    return false
}
