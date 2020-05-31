/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
   stack := []*Node{}
    res := []int{}
    if (root == nil) {
        return []int{}
    }
    stack = append(stack, root)

    for len(stack) > 0 {
        cur := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, cur.Val)

        for i := len(cur.Children) - 1; i >= 0; i-- {
            stack = append(stack, cur.Children[i])
         }
    }
    
    return res
}
