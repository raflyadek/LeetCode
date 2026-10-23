func returnToBoundaryCount(nums []int) int {
    step := 0
    res := 0
    for _, v := range nums {
        step += v
        if step == 0 {
            res++
        }
    }
    return res
}