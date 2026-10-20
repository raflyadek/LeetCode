func minOperations(nums []int, k int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < k {
            res++
        }
    }
    return res
}