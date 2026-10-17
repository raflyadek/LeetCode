func intersect(nums1 []int, nums2 []int) []int {
    result := make([]int,0)
    mapNum := make(map[int]int)
    for _, v := range nums1 {
        mapNum[v]++
    }
    for _, v := range nums2 {
        if _, found := mapNum[v]; found && mapNum[v] > 0 {
            mapNum[v]--
            result = append(result, v)
        }
    }

    return result
}