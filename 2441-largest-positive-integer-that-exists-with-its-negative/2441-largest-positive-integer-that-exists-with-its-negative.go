func findMaxK(nums []int) int {
    mapNum := make(map[int]bool)
    result := -1
    for _, v := range nums {
        mapNum[v] = true
    }

    for k, _ := range mapNum {
        if mapNum[-k] && k > result {
            result = k
        }
    }
    return result 
}