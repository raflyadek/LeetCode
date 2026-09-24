func singleNumber(nums []int) int {
    slices.Sort(nums)
    lastIndex := len(nums)-1
    for i := 0; i < len(nums); i++ {
        if i == lastIndex {
            return nums[lastIndex]
        }
        if nums[i] == nums[i+1] {
            i += 1
            continue
        } else {
            return nums[i]
        }
    }
    return 0
}