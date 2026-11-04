func isPowerOfFour(n int) bool {
    //same with power of 3
    if n == 1 {
        return true
    } else if n > 1 && n % 4 == 0 {
        return isPowerOfFour(n/4)
    } else {
        return false
    }
}