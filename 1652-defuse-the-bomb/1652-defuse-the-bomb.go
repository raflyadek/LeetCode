func decrypt(code []int, k int) []int {
	lenC := len(code)
    result := make([]int, lenC)
    if k == 0 {
        //because length of result is lenC then if nothing is append to this variable then its filled with zero-value of the data type
        return result
    }

    //index for the value we want to replace 
    for i := 0; i < len(code); i++ {
        //for the value we want to provide to the result[i]
        sum := 0
        if k > 0 {
            for j := 1; j <= k; j++ {
                sum += code[(i+j)%lenC]
            }
        } else {
            //we add lenC as the index because it will -, and lenC will acts as a jump to the last index instead of -, for example lenC = 4 0 - 1 + 4 is -1+4 = 3 -> 3 % 4 is 3 so it will snatch code[3] to the sum and we continue with j = 2 -> (0-2+4)%3 = 2 % 3 = 2 and so on untill j <= k 
            //-k for -(-k) so -(-2) = 2
            for j := 1; j <= -k; j++ {
                sum += code[(i-j+lenC)%lenC]
            }
        }
        //append the sum to result
        result[i] = sum 
    }

    return result
}
