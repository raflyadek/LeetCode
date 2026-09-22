func searchInsert(nums []int, target int) int {
    lastIndex := len(nums)-1
    if nums[lastIndex] < target {
        return lastIndex+1
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > target {
            if i == 0 {
                return 0
            }
            return i 
        }
        if nums[i] == target {
            return i
        }
    }
    return 0
}