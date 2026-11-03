func longestCommonPrefix(strs []string) string {
    if len(strs) <= 1 {
        return strs[0]
    }
    n := len(strs)-1
    sort.Strings(strs)
    for i := 0; i < len(strs[0]); i++ {
        //just compare first and last element
        //because it needs to match all the elements if 1 not have the prefix 
        //then its empty
        if strs[0][i] != strs[n][i] {
            return strs[0][:i]
        }
    }
    return strs[0]
}