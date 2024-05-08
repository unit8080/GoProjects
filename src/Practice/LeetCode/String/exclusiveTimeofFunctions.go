// 636. Exclusive Time of Functions
// https://leetcode.com/problems/exclusive-time-of-functions/description/

func exclusiveTime(n int, logs []string) []int {

	stack := make([][]int, 0)
	ans := make([]int, n)

	index := 0
	for index < len(logs) {

		splitStr := strings.Split(logs[index], ":")
		pid, _ := strconv.Atoi(splitStr[0])
		action := splitStr[1]
		time, _ := strconv.Atoi(splitStr[2])

		if action == "start" {
			stack = append(stack, []int{pid, time})
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			diff := time - top[1] + 1
			ans[top[0]] += diff

			if len(stack) > 0 {
				pidOnStack := stack[len(stack)-1][0]
				ans[pidOnStack] -= diff
			}
		}
		index++
	}
	return ans
}
