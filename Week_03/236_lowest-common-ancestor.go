/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 var parent *TreeNode
 func dfs(root, p, q *TreeNode) bool{
     if root == nil {
         return false
     }

     l := dfs(root.Left, p, q)
     r := dfs(root.Right, p, q)

     if (l && r) || (root.Val == p.Val || root.Val == q.Val) && (l || r) {
         parent = root
     }

     return (l || r || root.Val == p.Val || root.Val == q.Val)
 }

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent = nil
    dfs(root,p,q)

    return parent
}
