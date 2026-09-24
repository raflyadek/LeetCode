func containsDuplicate(nums []int) bool {
    mapNums := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        mapNums[nums[i]]++
        if mapNums[nums[i]] > 1 {
            return true
        }
    }
    return false
}