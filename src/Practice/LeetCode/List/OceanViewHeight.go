// 1762. Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/

func findBuildings(heights []int) []int {
    result := []int{}
    i := len(heights) - 1
    taller := heights[i]
    result = append(result, i)
    i--
    for i >= 0 {
        if heights[i] > taller {
            taller = heights[i]
            result = append(result, i)
        }
        i--
    }

    left := 0
    right := len(result) -1
    for left < right {
        result[left], result[right] = result[right], result[left]
        left++
        right--
    }

    return result
}
