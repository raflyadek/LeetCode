func canBeEqual(target []int, arr []int) bool {
    if len(target) != len(arr) {
        return false
    }
    //make a map and store it and then compare to the second parameter
    mapN := make(map[int]int)
    for i := 0; i < len(target); i++ {
        mapN[target[i]]++
        mapN[arr[i]]--
    }
    for _, v := range mapN {
        if v > 0 {
            return false
        }
    }

    return true
}