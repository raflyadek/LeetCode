func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    temp := x
    res := 0
    for x > 0 {
        s := x % 10
        res = res * 10 + s
        x /= 10
    }
    return temp == res
}