/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
     if root == nil {
        return [][]int{}
    } 

    res := [][]int{}
    que := []*Node{}
    que = append(que, root)
    for len(que) > 0 {
        size := len(que)
        level := []int{}
        for i:= 0; i < size; i++ {
            cur := que[0]
            level = append(level, cur.Val)
            que = que[1:len(que)]
            que = append(que, cur.Children...)
        }

        res = append(res, level)
    }
    
    return res
}
