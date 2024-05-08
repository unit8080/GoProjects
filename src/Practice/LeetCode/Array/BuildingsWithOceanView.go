// 1762. Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/

func findBuildings(heights []int) []int {
    result := []int{}
    i := len(heights) - 1
    taller := heights[i] // tallest so far
    result = append(result, i)
    i--
    for i >= 0 {
        if heights[i] > taller { // only taller can see
            taller = heights[i] // next taller
            result = append(result, i)
        }
        i--
    }

    // sorted order for indexes - since appending above
    left := 0
    right := len(result) -1
    for left < right {
        result[left], result[right] = result[right], result[left]
        left++
        right--
    }

    return result
}
