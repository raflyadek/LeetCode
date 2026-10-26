func sumOfSquares(nums []int) int {
    n := len(nums)
    res := 0
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            res += nums[i-1] * nums[i-1]
        }
    }   
    return res
}