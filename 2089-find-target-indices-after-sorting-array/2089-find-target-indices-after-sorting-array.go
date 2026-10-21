func targetIndices(nums []int, target int) []int {
    slices.Sort(nums)
    sliceRes := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            sliceRes = append(sliceRes, i)
        }
    }
    return sliceRes
}