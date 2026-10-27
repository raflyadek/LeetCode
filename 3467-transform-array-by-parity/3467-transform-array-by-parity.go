func transformArray(nums []int) []int {
    zero := 0
    n := len(nums)
    one := n-1
    res := make([]int, n)
    for _, v := range nums {
        if v % 2 == 0 {
            res[zero] = 0
            zero++
        } else {
            res[one] = 1
            one--
        }
    }
    return res
}