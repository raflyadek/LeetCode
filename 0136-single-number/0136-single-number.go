func singleNumber(nums []int) int {
    slices.Sort(nums)
    //edge case 
    if len(nums) == 1 {
        return nums[0]
    }
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            return nums[i-1]
        } else {
            i += 1
            if i == len(nums) - 1 {
                return nums[i]
            }
        }
    }
    return 0
}