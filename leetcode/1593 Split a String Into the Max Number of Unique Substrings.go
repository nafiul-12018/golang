/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(r *TreeNode, level int, dp *[]int) {

    if r== nil{
        return
    }

    if level == len(*dp) {
        *dp = append(*dp, r.Val)
    } else {
        (*dp)[level] += r.Val
    }

    dfs(r.Left, level+1, dp)
    dfs(r.Right, level+1, dp)
} 
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    dp := make([]int, 0)
    dfs(root, 0, &dp)
    //fmt.Println(dp)
    sort.Ints(dp)
    //fmt.Println(dp)
    n := len(dp)
    if n < k {
        return -1
    }
    return int64(dp[n-k])
}   
