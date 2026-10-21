func getSneakyNumbers(nums []int) []int {
    mapN := make(map[int]int)
    for _, v := range nums {
        mapN[v]++
    }
    res := make([]int, 0)
    for k, v := range mapN {
        if v == 2 {
            res = append(res, k)
        }
    }
    return res
}