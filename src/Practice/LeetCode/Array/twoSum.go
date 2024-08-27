// 1. Two Sum
// https://leetcode.com/problems/two-sum/

// func twoSum(nums []int, target int) []int {
//     // ret := []int {-1,-1}
//     if len(nums) < 2 {
//         return []int {-1,-1}
//     }

//     nMap := make(map[int]int)
//     for i, num := range nums {
//         num2 := target - num
//         if idx, ok := nMap[num2]; ok {
//             // ret = []int{idx, i}
//             // return ret
//             return []int{idx, i}

//         }
//         nMap[num] = i
//     }
//     return []int{-1, -1}
// }

// func twoSum(nums []int, target int) []int {

// 	if len(nums) < 2 {
// 		return []int{}
// 	}


//     targetMap := make(map[int]int, 0)
//     for idx, num := range nums {
//         targetMap[num] = idx
//     }
//     for idx1, num := range nums {
//         otherNum := target - num 
//         idx2, exists :=  targetMap[otherNum]
//         if exists && idx1 != idx2 {
//             return []int{idx1, idx2}
//         }
//     }
//     return []int{}
// }

func twoSum(nums []int, target int) []int {

    targetMap := make(map[int]int)
    for idx, num := range nums {
        idx2, ok := targetMap[target - num]
        if !ok {
            targetMap[num] = idx
        } else {
            return []int {idx, idx2}
        }
    }
    return []int{}
}
