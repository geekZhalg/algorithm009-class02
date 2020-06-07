func backtrace(nums []int, res[]int, result *[][]int) {
    if len(nums) == 0 {
        tmp := make([]int, len(res))
        copy(tmp, res)
        (*result) = append(*result, tmp)
        return 
    }

    for i := 0; i < len(nums); i++ {
        res = append(res, nums[i])
        tmp:= append([]int{}, nums[:i]...)
        tmp = append(tmp, nums[i+1:]...)
        //fmt.Println(tmp)
        backtrace(tmp, res, result)
        res = res[:len(res) - 1]
    }

    return
}

func permute(nums []int) [][]int {
    result := [][]int{}

    backtrace(nums, []int{}, &result)
    return result
}
