func findLHS(nums []int) int {
    mapNum := make(map[int]int)
    for _, v := range nums {
        mapNum[v]++
    }
    temp, high := 0, 0
    for k, v := range mapNum {        
        if val, found := mapNum[k+1]; found {
            temp += val + v
            if temp > high {
                high = temp
            }
            temp = 0
        }
    }
    return high
}