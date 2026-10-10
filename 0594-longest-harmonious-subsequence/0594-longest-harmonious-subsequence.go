func findLHS(nums []int) int {
    mapNum := make(map[int]int)
    for _, v := range nums {
        mapNum[v]++
    }

    tempK, counter, high := 0, 0, 0
    for k, v := range mapNum {
        tempK = k
        counter = v
        if value, found := mapNum[tempK+1]; found {
            counter += value
            if counter > high {
                high = counter
            }
            counter = 0
        } else {
            counter = 0
        }
    }
    return high
}