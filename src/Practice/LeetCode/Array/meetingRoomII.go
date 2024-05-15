// 253. Meeting Rooms II
// https://leetcode.com/problems/meeting-rooms-ii/

func minMeetingRooms(intervals [][]int) int {

    n := len(intervals)
    start := make([]int, n)
    end := make([]int, n)
    
    for i, interval := range intervals {
        start[i] = interval[0]
        end[i] = interval[1]
    }
    sort.Ints(start)
    sort.Ints(end)

    room := 0
    s_ptr := 0
    e_ptr := 0
    for s_ptr < n {
        if start[s_ptr] < end[e_ptr] {
            room++
        } else {
            e_ptr++
        }
        s_ptr++
    }
    return room
}
