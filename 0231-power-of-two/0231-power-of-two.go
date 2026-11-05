func isPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    }
    for n > 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            return false
        }
    }
    return n == 1
}