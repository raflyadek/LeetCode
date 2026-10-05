func missingNumber(nums []int) int {
    slices.Sort(nums)
    if len(nums) == 1 {
        if nums[0] == 0 {
            return 1
        } else if nums[0] ==1 {
            return 0
        }
    }
    if nums[0] != 0 {
        return 0
    }
    for i := 1; i < len(nums); i++ {
        if nums[i]-1 != nums[i-1] {
            return i
        } 
        if i == len(nums)-1 {
            return i+1
        }
    }
    return 0
}