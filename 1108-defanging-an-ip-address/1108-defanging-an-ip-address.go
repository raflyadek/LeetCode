func defangIPaddr(address string) string {
    res := ""
    for i := 0; i < len(address); i++ {
        if string(address[i]) == "." {
            res += fmt.Sprintf("[%s]", string(address[i]))
        } else {
            res += string(address[i])
        }
    }
    return res
}