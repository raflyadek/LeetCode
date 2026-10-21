func concatWithReverse(nums []int) []int {
    n := len(nums)
    newNums := make([]int, n *2)
    for i := 0; i < len(nums); i++ {
        newNums[i] = nums[i]
        newNums[i+n]=nums[n-i-1]
    }
    return newNums
}