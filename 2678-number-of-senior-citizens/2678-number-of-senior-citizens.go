func countSeniors(details []string) int {
    res := 0
    for i := 0; i < len(details); i++ {
        num, _ := strconv.Atoi(details[i][11:13])
        if num > 60 {
            res++
        }
    }
    return res
}