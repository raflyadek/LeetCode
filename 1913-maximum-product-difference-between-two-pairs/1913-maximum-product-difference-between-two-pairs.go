func maxProductDifference(nums []int) int {
    n := len(nums)
    slices.Sort(nums)

    return (nums[n-1]*nums[n-2]) - (nums[0]*nums[1])
}