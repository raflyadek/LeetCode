func majorityElement(nums []int) int {
    half := len(nums)/2
    mapNum := make(map[int]int)
    for _, v := range nums {
        mapNum[v]++
        if mapNum[v] > half {
            return v
        }
    }
    return 0
}