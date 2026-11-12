func checkGoodInteger(n int) bool {
    square := 0
    digit := 0
    for n > 0 {
        s := n % 10
        digit += s
        square += s * s
        n /= 10
    }
    return (square-digit) >= 50
}