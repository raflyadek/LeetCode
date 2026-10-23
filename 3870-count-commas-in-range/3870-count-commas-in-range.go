func countCommas(n int) int {
    if n < 1000 {
        return 0
    }
    res := n - 1000

    return res+1
}