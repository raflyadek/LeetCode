func isMiddleElementUnique(nums []int) bool {
    mapN := make(map[int]int)
    half := len(nums)/2
    mid := 0
    for i, v := range nums {
        mapN[v]++
        if i == half {
            mid = v
        }
    }
    val, _ := mapN[mid]
    if val > 1 {
        return false
    }
    return true
}