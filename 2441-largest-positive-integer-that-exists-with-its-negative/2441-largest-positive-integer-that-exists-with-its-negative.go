func findMaxK(nums []int) int {
    mapNum := make(map[int]bool)
    result := -1
    for _, v := range nums {
        mapNum[v] = true
    }

    for k, _ := range mapNum {
        if k > 0 && k > result {
            if mapNum[-k] {
                result = k
            }
        }
    }
    return result 
}