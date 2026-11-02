func calPoints(operations []string) int {
    //+ is plus last index and index before that 
    //D is *2
    //C is remove the last index
    res := 0
    temp := make([]int, 0)
    for _, v := range operations {
        lastIndex := len(temp)-1
        if v == "D" {
            n := temp[lastIndex] * 2
            temp = append(temp, n)
            continue
        }
        if v == "C" {
            temp = temp[:lastIndex]
            continue
        }
        if v == "+" {
            n := temp[lastIndex] + temp[lastIndex-1]
            temp = append(temp, n)
        }
        if num, err := strconv.Atoi(v); err == nil {
            temp = append(temp, num)
        }
    }
    for _, v := range temp {
        res += v
    }
    return res
}