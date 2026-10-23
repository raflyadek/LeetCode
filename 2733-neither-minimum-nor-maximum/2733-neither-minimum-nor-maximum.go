func findNonMinOrMax(nums []int) int {
    if len(nums) == 2 {
        return -1
    }
    min, max := nums[0], nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        } else if v < min {
            min = v
        }
    }
    for _, v := range nums {
        if v != max && v != min {
            return v
        }
    }
    return -1
}