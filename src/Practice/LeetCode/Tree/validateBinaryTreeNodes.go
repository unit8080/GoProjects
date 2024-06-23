// 1361. Validate Binary Tree Nodes
// https://leetcode.com/problems/validate-binary-tree-nodes/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    
    hasParent := make(map[int]struct{})
    for _, child := range leftChild {
        if child != -1 {
            hasParent[child] = struct{}{}
        }
    }
    for _, child := range rightChild {
        if child != -1 {
            hasParent[child] = struct{}{}
        }
    }
    if len(hasParent) == n {
        return false
    }
    root := -1
    for i := 0; i < n; i++ {
        if _, exists := hasParent[i]; !exists {
            root = i
            break
        }
    }
    visited := map[int]bool{}
    var dfs func(root int) bool 
    dfs = func(root int) bool {
        if root == -1 {
            return true
        }
        if visited[root] {
            return false
        }
        visited[root] = true
        return dfs(leftChild[root]) && dfs(rightChild[root])
    }
    return dfs(root) && len(visited) == n
}
