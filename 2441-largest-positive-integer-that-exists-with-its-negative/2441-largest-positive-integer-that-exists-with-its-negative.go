func findMaxK(nums []int) int {
    slices.Sort(nums)
    mapNum := make(map[int]bool)
    result := -1
    for _, v := range nums {
        mapNum[v] = true
        if mapNum[-v] && v > result {
            result = v
        }
    }

    return result 
}