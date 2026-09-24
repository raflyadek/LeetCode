func singleNumber(nums []int) int {
    mapNum := make(map[int]int) 
    for i := 0; i < len(nums); i++ {
        mapNum[nums[i]]++
    }
    
    for k, v := range mapNum {
        if v == 1 {
            return k
        }
    }
    return 0
}