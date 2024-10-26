/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, cnt map[int]int) int {
    if root ==  nil {
        return 0
    }

    if val, exists := cnt[root.Val] ; exists == true {
        return val
    }

    mx := 0
    left := dfs(root.Left, cnt)
    if left >= mx {
        mx = left
    }

    right := dfs(root.Right, cnt)
    
    if right >= mx {
        mx = right
    }

    cnt[root.Val] = mx+1
    return mx+1
} 

func dfs2(root *TreeNode, cur, up int, pre map[int]int, cost map[int]int) {
    if root ==  nil {
        return 
    }

    cost[root.Val] = up

    left := dfs(root.Left, pre)
    rightMx := left + cur +1
    if rightMx <= up {
        rightMx = up
    }
    dfs2(root.Right, cur+1, rightMx, pre, cost)
    right := dfs(root.Right, pre)
    leftMx := right + cur +1

    if leftMx <= up {
        leftMx = up 
    }

    dfs2(root.Left, cur+1, leftMx, pre, cost)
} 

func treeQueries(root *TreeNode, queries []int) []int {
    var pathCnt map[int]int
    pathCnt = make(map[int]int)
    cost := make(map[int]int)

    dfs(root, pathCnt) 
    dfs2(root,0,0, pathCnt, cost)
    var ans []int

    for _, val := range queries {
        ans = append(ans, cost[val]-1)
    }
    return ans
}
