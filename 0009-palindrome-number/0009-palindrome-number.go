func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    temp := x
    res := 0
    for x > 0 {
        res = res * 10 + (x%10)
        x /= 10
    }
    return res == temp 
}