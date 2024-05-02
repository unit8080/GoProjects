// 973. K Closest Points to Origin
// https://leetcode.com/problems/k-closest-points-to-origin/

func kClosest(points [][]int, k int) [][]int {

    distance := func(point []int) int {
        return point[0] * point[0] + point[1] * point[1]
    }
    distances := make([]int, len(points))
    for i, point := range points {
        distances[i] = distance(point)
    }
    
    partition := func(l, r int) int {
        for l < r {
            if distances[l+1] <= distances[l] {
                points[l+1], points[l], distances[l+1], distances[l] = points[l], points[l+1], distances[l], distances[l+1]
                l++
            } else if distances[r] > distances[l] {
                r--
            } else {
                points[r], points[l+1], distances[r], distances[l+1] = points[l+1], points[r], distances[l+1], distances[r]
            }
        }
        return l
    }
    var qs func(l, r int) [][]int

    qs = func(l, r int) [][]int {
        p := partition(l, r)
        if k-1  < p {
            return qs(0, p-1)
        } else if k-1 > p {
            return qs(p+1, r)
        } else {
            return points[:p+1]
        }
    }
    return qs(0, len(points) -1)
}
