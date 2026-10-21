func getSneakyNumbers(nums []int) []int {
    mapN := make(map[int]bool)
    res := make([]int, 0)
    for _, v := range nums {
        if mapN[v] {
            res = append(res, v)
        }
        mapN[v] = true
    }
    
    return res
}