// 210. Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/?envType=company&envId=nvidia&favoriteSlug=nvidia-thirty-days

func findOrder(numCourses int, prerequisites [][]int) []int {
    
    n := len(prerequisites)
    countZero := 0 
    indegree := make([]int, numCourses)
    adjList := make(map[int][]int, 0)
    for i := 0; i < n; i++ {
        u := prerequisites[i][0]
        v := prerequisites[i][1]
        adjList[u] = append(adjList[u], v)
        indegree[v] += 1
    }

    result := []int{}
    queue := []int{}
    for u := 0; u < numCourses; u++ {
        if indegree[u] == 0 {
            result = append(result, u)
            countZero++
            queue = append(queue, u)
        }
    }

    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        // list = append(list, course)
        courses, exists := adjList[course]
        if !exists {
            continue
        }
        for _, course := range courses {
            indegree[course]--
            if indegree[course] == 0 {
                result = append(result, course)
                countZero++
                queue = append(queue, course)
            }
        }
    }
    if countZero == numCourses {
        l := 0
        r := len(result)-1
        for l < r {
            result[l], result[r] = result[r], result[l]
            l++
            r--
        }
        return result
    } else {
        return []int{}
    }
}
