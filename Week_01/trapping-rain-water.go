
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func trap(height []int) int {
    var stack []int
    var res int = 0
    var i  int
    for i = 0; i < len(height); i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack) - 1]] {
            //stack_top = stack[len(stack) - 1]
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]//pop

            if len(stack) == 0 {
                break
            }

            wide := i - stack[len(stack) - 1] - 1
            h := min(height[i], height[stack[len(stack) - 1]]) - height[top]
            //fmt.Println(wide, h)
            res += wide * h
        }

        stack = append(stack, i)
    }

    return res
}
