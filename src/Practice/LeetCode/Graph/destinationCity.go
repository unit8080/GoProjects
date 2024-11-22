// 1436. Destination City
// https://leetcode.com/problems/destination-city/

func destCity(paths [][]string) string {
    
    city := map[string]bool{}

    for i := 0; i < len(paths); i++ {
        path := paths[i]
        city[path[0]] = true
    }
    for i := 0; i < len(paths); i++ {
        path := paths[i]
        if _, exists := city[path[1]]; !exists {
            return path[1]
        }
    }
    return ""
}
