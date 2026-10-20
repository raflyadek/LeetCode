func minimumOperations(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] % 3 != 0 {
            if (nums[i] - 1)%3 == 0 || (nums[i] + 1)% 3 == 0 {
                res++
            }
        }
    }
    return res
}