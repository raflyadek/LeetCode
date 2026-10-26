func canBeEqual(target []int, arr []int) bool {
    //make a map and store it and then compare to the second parameter
    mapN := make(map[int]int)
    for _, v := range target {
        mapN[v]++
    }
    for _, v := range arr {
        mapN[v]--
    }
    for _, v := range mapN {
        if v > 0 {
            return false
        }
    }

    return true
}