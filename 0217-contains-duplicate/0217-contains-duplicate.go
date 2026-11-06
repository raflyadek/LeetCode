func containsDuplicate(nums []int) bool {
    res := make(map[int]bool)
    for _, v := range nums {
        if res[v] {
            return true
        }
        res[v] = true
    }
    return false
}