func missingNumber(nums []int) int {
    slices.Sort(nums)
    n := len(nums)
    if nums[n-1] != n {
        return n
    }
    for i := 0; i < len(nums); i++ {
        if i != nums[i] {
            return i
        }
    }
    return 0
}