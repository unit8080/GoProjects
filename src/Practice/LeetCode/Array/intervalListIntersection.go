// 986. Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    
    ans := [][]int{}

    i := 0
    j := 0
    len1 := len(firstList)
    len2 := len(secondList)

    for i < len1 && j < len2 {

        if firstList[i][1] >= secondList[j][0] && firstList[i][0] <= secondList[j][1] {
            low := max(firstList[i][0], secondList[j][0])
            high := min(firstList[i][1], secondList[j][1])
            ans = append(ans, []int{low, high})
        }
        if firstList[i][1] < secondList[j][1] {
            i++
        } else {
            j++
        }
    }
    
    return ans
}
