/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, level int, depth *[]int) {
    if root == nil {
        return 
    }

    root.Val = (*depth)[level] - root.Val
    dfs(root.Left, level +1, depth)
    dfs(root.Right, level +1, depth)
     
}

func cal(root *TreeNode, level int, depth *[]int) {
    if root == nil {
        return 
    }

    if len(*depth) == level {
        *depth = append(*depth, root.Val)
    } else {
        (*depth)[level] += root.Val
    }


    cal(root.Left, level +1, depth)
    cal(root.Right, level +1, depth)

    if root.Left != nil && root.Right != nil {
        sum := root.Left.Val + root.Right.Val
        root.Left.Val = sum
        root.Right.Val = sum
    }
    //root.Val = x+y
    return 
}


func replaceValueInTree(root *TreeNode) *TreeNode {
    var depth []int
    cal(root, 0, &depth)

    // for _,item:= range depth {
    //     fmt.Println(item)
    // }

    dfs(root, 0, &depth)

    return root
}
