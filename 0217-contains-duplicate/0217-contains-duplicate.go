func containsDuplicate(nums []int) bool {
    result := false
    numberMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        numberMap[nums[i]]++
    }
    for _, v  := range numberMap {
        if v > 1 {
            result = true
        } 
    }
    return result
}