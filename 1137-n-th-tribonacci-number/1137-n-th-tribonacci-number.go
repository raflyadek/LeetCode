func tribonacci(n int) int {
    res := make([]int, 0, n)
    res = append(res, 0, 1, 1)
    for i := 3; i <= n; i++ {
        sum := res[i-3] + res[i-2] + res[i-1]
        res = append(res, sum)
    }
    return res[n]
}