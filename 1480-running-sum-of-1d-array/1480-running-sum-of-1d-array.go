func runningSum(nums []int) []int {
    numsZ := make([]int, len(nums))
    numsZ[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        numsZ[i] = numsZ[i-1] + nums[i]
    }
    return numsZ
}