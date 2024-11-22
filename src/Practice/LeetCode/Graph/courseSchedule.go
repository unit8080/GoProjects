// 207. Course Schedule
// https://leetcode.com/problems/course-schedule/

func canFinish(numCourses int, prerequisites [][]int) bool {

	adj := map[int][]int{}
	n := numCourses
	indegree := make([]int, n)

    if len(prerequisites) == 0 {
        return true
    }
	for i := 0; i < len(prerequisites); i++ {
		u := prerequisites[i][0]
		v := prerequisites[i][1]
		adj[u] = append(adj[u], v)
		indegree[v]++
	}

	count := 0
	queue := []int{}

	for i := 0; i < n; i++ {
		
		if indegree[i] == 0 {
			count++
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		courses, exists := adj[course]
        if !exists {
            continue
        }
		for _, v := range courses {
			indegree[v]--
			if indegree[v] == 0 {
				count++
				queue = append(queue, v)
			}
		}
	}
	if count == n {
		return true
	} else {
		return false
	}
}
