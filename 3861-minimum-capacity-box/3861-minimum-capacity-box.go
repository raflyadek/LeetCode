func minimumIndex(capacity []int, itemSize int) int {
    res := -1
    flag := false
    temp := 0
    for i := 0; i < len(capacity); i++ {
        if capacity[i] >= itemSize {
            if flag && (capacity[i] > temp || capacity[i] == temp) {
                continue
            }
            temp = capacity[i]
            res = i
            flag = true
        }
    }
    return res
}