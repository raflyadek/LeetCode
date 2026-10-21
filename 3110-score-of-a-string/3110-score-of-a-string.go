func scoreOfString(s string) int {
    sum := 0
    left, right := 0, 1
    for right < len(s) {
        sum += abs(int(s[left])-int(s[right]))
        left++
        right++
    }
    return sum
}
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}