func alternateDigitSum(n int) int {
    c := 0
    res := 0
    rev := 0
    for n > 0 {
        s := n % 10
        rev = (rev * 10) + s
        n /= 10
    }
    for rev > 0 {
        s := rev % 10 
        if c % 2 == 0 {
            res += s
        } else {
            res -= s
        }
        rev /= 10
        c++
    }
    return res
}