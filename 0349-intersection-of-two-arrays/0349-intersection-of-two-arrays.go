func intersection(nums1 []int, nums2 []int) []int {
    mapNum := make(map[int]bool)
    res := make([]int,0)
    for i := 0; i <len(nums1); i++ {
        mapNum[nums1[i]] = true
    }
    for i := 0; i <len(nums2); i++ {
        if mapNum[nums2[i]] {
            mapNum[nums2[i]] = false
            res = append(res, nums2[i])
        }
    }
    return res
}