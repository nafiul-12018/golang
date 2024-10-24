func dfs(r1 *TreeNode, r2 *TreeNode) bool {
    
    if r1 == nil {
        return r2 == nil
    }
    if r2 == nil {
        return false
    }
    //fmt.Println(r1.Val, " ", r2.Val)
    if r1.Val != r2.Val {
        return false
    }

   
    
    ret := (dfs(r1.Left, r2.Right) && dfs(r1.Right, r2.Left)) || (dfs(r1.Left, r2.Left) && dfs(r1.Right, r2.Right))
    
    return ret
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    return dfs(root1, root2)
}
