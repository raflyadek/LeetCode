func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        lowest := nums[0]
        index := 0 
        for j := 0; j < len(nums); j++ {
            if nums[j] < lowest {
                lowest = nums[j]
                index = j
            }
        }
        nums[index] = lowest * multiplier
    }
    return nums
}