func numIdenticalPairs(nums []int) int {
    left, right := 0, 1
    res := 0
    for left < len(nums) {
        if nums[left] == nums[right] {
            res++
        } 
        right++
        if right == len(nums) {
            left += 1
            right = left+1
            if right == len(nums) {
                break
            }
        }
    }
    return res
}