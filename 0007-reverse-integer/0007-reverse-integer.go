func reverse(x int) int {
    //is it cheating using atoi/itoa???hm
    str := strconv.Itoa(x)
    num := 0
    limitPstv := 2 << 30
    limitNgtv := -2 << 30
    if x < 0 {
        numtemp, _ := strconv.Atoi(str[1:])
        for numtemp > 0 {
            s := numtemp%10
            num = num * 10 + s
            numtemp /= 10
        }
        if -num < limitNgtv {
            return 0
        }
        return -num
    }
    for x > 0 {
        s := x % 10
        num = num * 10 + s
        x /= 10
    }
    if num > limitPstv {
        return 0
    }
    return num
}