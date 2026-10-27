func toggleLightBulbs(bulbs []int) []int {
    mapB := make(map[int]int)
    for _, v := range bulbs {
        mapB[v]++
    }
    res := make([]int, 0)
    for k, v := range mapB {
        if v % 2 != 0 {
            res = append(res, k)
        }
    }
    slices.Sort(res)
    return res
}