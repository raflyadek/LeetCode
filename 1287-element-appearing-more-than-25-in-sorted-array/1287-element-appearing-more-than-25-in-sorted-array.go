func findSpecialInteger(arr []int) int {
    max := len(arr)/4
    mapN := make(map[int]int)
    for _, v := range arr {
        mapN[v]++
        if mapN[v] > max {
            return v
        }
    }
    return 0
}