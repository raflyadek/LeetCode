func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    mapElement := make(map[int]int)
    result := 0
    resultTemp := 0
    for i := 0; i < len(nums); i++ {
        mapElement[nums[i]]++
    }

    for k, v := range mapElement {
        if resultTemp < v {
            result = k   
            resultTemp = v
        }
    }
    return result
}