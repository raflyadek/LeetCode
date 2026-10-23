func maxContainers(n int, w int, maxWeight int) int {
    cell := n * n
    t := w
    res := 0
    for w <= maxWeight && cell != 0 {
        w += t
        res++
        cell--
    }
    return res
}