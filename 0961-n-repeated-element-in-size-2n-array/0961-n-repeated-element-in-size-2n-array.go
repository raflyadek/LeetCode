func repeatedNTimes(nums []int) int {
    mapN := make(map[int]int)
    for _, v := range nums {
        mapN[v]++
        if mapN[v] >= 2 {
            return v
        }
    }
    return 0
}