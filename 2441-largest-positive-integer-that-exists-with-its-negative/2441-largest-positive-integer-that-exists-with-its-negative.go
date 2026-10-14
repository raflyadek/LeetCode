func findMaxK(nums []int) int {
    result := 0
    slices.Sort(nums)
    mapN := make(map[int]bool)
    for i := len(nums)-1; i >= 0; i-- {
        mapN[nums[i]] = true
        if mapN[nums[i]*-1] {
            if nums[i]*-1 > result {
                result = nums[i]*-1
            }
        }
    }
    if result == 0 {
        result = -1
    }
    return result
}