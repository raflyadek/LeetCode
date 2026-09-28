func intersection(nums1 []int, nums2 []int) []int {
    result := make([]int,0,0)
    map1 := make(map[int]int)
    map2 := make(map[int]int)
    for i := 0; i < len(nums1); i++ {
        map1[nums1[i]]++
    }
    for j := 0; j < len(nums2); j++ {
        map2[nums2[j]]++
    }

    for k, _ := range map1 {
        if _, found := map2[k]; found {
            result = append(result, k)
        }
    }
    return result
}