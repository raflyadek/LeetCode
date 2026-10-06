func twoSum(nums []int, target int) []int {
    mapNum := make(map[int]int)
    for i, val := range nums {
        diff := target - val
        if idx, found := mapNum[diff]; found {
            return []int{i, idx}
        }
        mapNum[val] = i
    }
    return nil
}