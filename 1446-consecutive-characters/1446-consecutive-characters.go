func maxPower(s string) int {
    if len(s) == 0 {
        return 0
    }
    if len(s) == 1 {
        return 1
    }
    before := s[0]
    count := 0
    temp := 0
    for i := 0; i < len(s); i++ {
        if before == s[i] {
            temp++
        } else {
            if temp > count {
                count = temp
                temp = 0
            }
            temp = 0
            before = s[i]
            temp++
        }
    }
    if temp > count {
        count = temp
    }
    return count 
}