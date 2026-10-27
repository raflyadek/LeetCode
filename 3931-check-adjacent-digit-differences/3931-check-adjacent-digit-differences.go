func isAdjacentDiffAtMostTwo(s string) bool {
    left, right := 0, 1
    for right < len(s) {
        num, _ := strconv.Atoi(string(s[left]))
        num1, _ := strconv.Atoi(string(s[right]))
        if abs(num-num1) > 2 { 
            return false
        } else {
            left = right
            right += 1
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}