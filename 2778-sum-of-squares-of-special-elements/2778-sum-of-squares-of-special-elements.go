func sumOfSquares(nums []int) int {
    n := len(nums)
    res := 0
    for i := range nums {
        if n%(i+1) ==0 {
            res += nums[i] * nums[i]
        }
    }
    return res
}