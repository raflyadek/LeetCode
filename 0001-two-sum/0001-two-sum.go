func twoSum(nums []int, target int) []int {
    result := make([]int, 0, 0)
    j := len(nums)-1
    for i := 0; i < len(nums); i++ {
        if i == j && j > 0 {
            i = 0
            j--
        }
        if nums[i] + nums[j] == target {
            result = append(result, i, j)
            break
        }

    }
    return result
}