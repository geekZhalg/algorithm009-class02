学习笔记

~
/* 冒泡 */
/*
func sortArray(nums []int) []int {
    for i := 0; i < len(nums) - 1;i++ {
        for j := 0; j < len(nums) - 1 - i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }

    return nums
}
*/

/* 插入 */
/*
func sortArray(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        prev := i - 1

        cur := nums[i]
        for prev >= 0 && cur < nums[prev] {
            nums[prev + 1] = nums[prev]
            prev--
        }

        //fmt.Println(prev, nums)


        nums[prev+1] = cur
    }

    return nums
}
*/

/*选择排序*/
/*
func sortArray(nums []int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        minidx := i
        for j := i + 1; j < len(nums);j++ {
            if nums[j] < nums[minidx] {
                minidx = j
            }
        }

        nums[i], nums[minidx] = nums[minidx], nums[i]
    }

    return nums
}
*/

/* 快排 */
/*
func parttional(nums[]int, start, end int) int{
    idx := end
    cnt := start
    for i := start; i < end; i++ {
        if nums[i] < nums[idx] {
            nums[cnt], nums[i] = nums[i], nums[cnt]
            cnt++
        }
    }

    nums[cnt], nums[idx] = nums[idx], nums[cnt]
    return cnt
}

func Qsort(nums []int, start, end int) {
    if start >= end {
        return
    }


    mid := parttional(nums, start, end)
    //fmt.Println(nums)
    Qsort(nums, start, mid - 1)
    Qsort(nums, mid+ 1, end)
}

func sortArray(nums []int)[]int {
    Qsort(nums, 0, len(nums) - 1)

    return nums
}
*/

/* 归并排序 */
func merge(nums[]int, start, mid, end int) {
    i := start
    j := mid + 1
    k := 0
    tmparray := make([]int, end - start + 1)
    //copy(tmparray,nums[start:end+1])

    for i <= mid && j <= end {
        if nums[i] <= nums[j] {
            tmparray[k] = nums[i]
            i++
        } else {
            tmparray[k] = nums[j]
            j++
        }

        k++
    }

    for  i <= mid {
        tmparray[k] = nums[i]
        i++
        k++
    }

    for j <= end {
        tmparray[k] = nums[j]
        j++
        k++
    }

    for l := 0; l < end - start + 1; l++ {
        nums[start + l] = tmparray[l]
    }
}

func mergeSort(nums []int, start, end int) {
    //fmt.Println(start, end)
    if start >= end {
        return
    }

    mid := start + ((end - start) >> 1)
    mergeSort(nums, start, mid)
    mergeSort(nums, mid + 1, end)

    merge(nums, start, mid, end)
}

func sortArray(nums []int)[]int {
    mergeSort(nums, 0, len(nums) - 1)

    return nums
}
