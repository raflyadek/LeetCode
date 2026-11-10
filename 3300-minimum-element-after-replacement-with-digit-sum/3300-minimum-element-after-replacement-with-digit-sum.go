func minElement(nums []int) int {
    arr := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        temp := 0
        for nums[i] > 0 {
            temp += nums[i] % 10
            nums[i] /= 10
        }
        arr = append(arr, temp)
    }
    res := arr[0]
    for _, v := range arr {
        if v < res {
            res = v
        }
    }
    return res
}