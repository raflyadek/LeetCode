func removeTrailingZeros(num string) string {
    res := len(num)
    for i := res-1; i > 0; i-- {
        if num[i] == 48 {
            res--
        } else {
            break
        }
    }
    return num[0:res]
}