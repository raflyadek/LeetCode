func isArraySpecial(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
    left := 0
    //left will be right and right will increment 1
    for right := 1; right < len(nums); right++ {
        if nums[left] % 2 == 0 && nums[right] % 2 == 0 {
            return false
        }
        if nums[left] % 2 != 0 && nums[right] % 2 != 0 {
            return false
        }
        left = right
    }
    return true
}