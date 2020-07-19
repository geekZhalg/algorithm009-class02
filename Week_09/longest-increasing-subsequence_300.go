func max(x, y int) int{
    if x > y {
        return x
    }
    return y
}

//dp[i] = max(dp[0]...dp[i-1]) + 1  //nums[i] > nums[.]
func lengthOfLIS(nums []int) int {

    if len(nums) == 0 {
        return 0
    }
    
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp [i] = 1
    }

    res := 1
    for i := 1; i < len(nums); i++ {
        maxl := 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                maxl = max(dp[j], maxl)
            }
        }

        dp[i] = maxl + 1
        res = max(dp[i], res)
    }

    return res
}
