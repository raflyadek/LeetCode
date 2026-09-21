func twoSum(nums []int, target int) []int {
    result := make([]int, 0, 0)
    j := len(nums)-1
    for i := 0; i < len(nums); i++ {
        if i == j && j > 0 {
            j -= 1
            i = 0
        }
        total := nums[i] + nums[j]
        if total == target {
            result = append(result, i, j)
            return result
        }
    }
    return []int{}
}