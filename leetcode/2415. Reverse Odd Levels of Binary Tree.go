/**


 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    qu := make([]*TreeNode, 0)

    depth := 0

    qu = append(qu, root)

    for len(qu) > 0 {
        n := len(qu)
        if depth % 2 == 1 {
            i := 0
            j := n-1
            for i<j {
                qu[i].Val, qu[j].Val = qu[j].Val, qu[i].Val
                i++
                j--
            }
        }

        for j := 0; j <n ; j++ {
            //fmt.Printf("%d ", qu[j].Val)
            if qu[j].Left != nil {
                qu = append(qu, qu[j].Left)
            }

            if qu[j].Right != nil {
                qu = append(qu, qu[j].Right)
            }
        }

        //fmt.Println(". done")

        qu = qu[n:]
        depth++
    }
    return root
}
