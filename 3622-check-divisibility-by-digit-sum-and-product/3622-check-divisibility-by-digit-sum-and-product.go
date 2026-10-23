func checkDivisibility(n int) bool {
    t := n
    mul := 1
    sum := 0
    for t > 0 {
        s := t % 10
        sum += s
        mul *= s
        t /= 10
    }
    if n % (sum+mul) == 0 {
        return true
    }

    return false
}