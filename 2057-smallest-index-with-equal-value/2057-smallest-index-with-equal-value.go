func smallestEqual(nums []int) int {
    //map
    res := -1
    for i, v := range nums {
        if i % 10 == v {
            res = i
            return res
        }
    }
    return res
}