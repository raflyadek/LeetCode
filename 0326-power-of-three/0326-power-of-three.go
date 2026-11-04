func isPowerOfThree(n int) bool {
    if n < 1 {
        return false
    }

    //modulo 3, then divide it by 3, if later it 1, then its power of 3
    //15 % 3 = 5, 5/3 != 1 <- false
    for n % 3 == 0 {
        n /= 3
    }
    return n == 1
}