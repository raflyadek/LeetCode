func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := make([]int, 0)
    for _, v := range nums1 {
        merged = append(merged, v)
    }

    for _, v := range nums2 {
        merged = append(merged, v)
    }
    var res float64
    slices.Sort(merged)
    mid := len(merged)/2 - 1
    if len(merged) % 2 == 0 {
        res += (float64(merged[mid]) + float64(merged[mid+1]) )/ 2
    } else {
        res += float64(merged[mid+1])
    }
    return res
}