func findLucky(arr []int) int {
    mapZ := make(map[int]int)
    result := -1
    for _, v := range arr {
        mapZ[v]++
    }
    for k, v := range mapZ {
        if k > result && k == v {
            result = k
        }
    }
    return result 
}