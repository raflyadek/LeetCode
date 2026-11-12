func isMiddleElementUnique(nums []int) bool {
    half := nums[len(nums)/2]
    for i, v := range nums {
        if i == len(nums)/2 {
            continue
        }
        if v == half {
            return false
        }
    }
    return true
}