func isPowerOfFour(n int) bool {
    //same with power of 3
    if n < 1 {
        return false
    }

    for n % 4 == 0 {
        n /= 4
    }
    return n == 1
}