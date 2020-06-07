/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type buildTreeCtx struct {
     m map[int]int
 }
func (g buildTreeCtx)buildt(preorder []int, pstart , pend int, inorder []int, istart, iend int) *TreeNode {
    if pstart >= pend {
        return nil
    }

    root := preorder[pstart]
    i := g.m[root]
    num := i - istart
    node := &TreeNode{Val:root}
    node.Left = g.buildt(preorder, pstart + 1, pstart + num + 1, inorder, istart, i)
    node.Right = g.buildt(preorder, pstart + num + 1, pend, inorder, i+1, iend)
    return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    g := buildTreeCtx{map[int]int{}}
    for i,v:=range inorder {
        g.m[v] = i
    }


    return g.buildt(preorder, 0, len(preorder), inorder, 0, len(inorder))
}
