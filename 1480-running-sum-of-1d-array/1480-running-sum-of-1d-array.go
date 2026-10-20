func runningSum(nums []int) []int {
    numsZ := make([]int, len(nums))
    right := 0
    for right < len(nums) {
        left, sum := 0, 0
        for left <= right {
            sum += nums[left]
            numsZ[right] = sum
            left++
        }
        right++
    }
    return numsZ
}