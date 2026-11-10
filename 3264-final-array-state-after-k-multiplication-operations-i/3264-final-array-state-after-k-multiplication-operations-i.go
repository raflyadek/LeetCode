func getFinalState(nums []int, k int, multiplier int) []int {
    res := nums
    for i := 0; i < k; i++ {
        lowest := nums[0]
        index := 0 
        for j := 0; j < len(nums); j++ {
            if nums[j] < lowest {
                lowest = nums[j]
                index = j
            }
        }
        res[index] = lowest * multiplier
    }
    return res
}