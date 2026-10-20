func subtractProductAndSum(n int) int {
    multiplyDigit := 1
    sumDigit := 0
    for n > 0 {
        s := n % 10
        multiplyDigit *= s
        sumDigit += s
        n /= 10
    }
    return multiplyDigit - sumDigit
}