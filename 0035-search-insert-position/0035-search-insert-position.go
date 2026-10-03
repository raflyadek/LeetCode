func searchInsert(nums []int, target int) int {
    //edge case
    if nums[len(nums)-1] < target {
        return len(nums)
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] > target || nums[i] == target {
            return i
        }
    }
    return 0
}