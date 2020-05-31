func twoSum(nums []int, target int) []int {
    var result []int
    mnums := make(map[int]int, 0)

    for i := range nums {
        if j,ok := mnums[target - nums[i]];ok {
            if i != j {
                result = append(result, i)
                result = append(result, j)
                break
            }   
        }

        mnums[nums[i]] = i
    }

/*
    for i,num := range nums {   
        if j,ok := mnums[target - num];ok {
            if i != j {
                result = append(result, i)
                result = append(result, j)
                break
            }   
        }
    }
*/
    return result
}
