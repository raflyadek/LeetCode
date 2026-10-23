func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        res[i] = nums[i] * nums[i]
    }
    slices.Sort(res)
    return res
}