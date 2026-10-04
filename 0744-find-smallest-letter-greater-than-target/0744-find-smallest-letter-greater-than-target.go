func nextGreatestLetter(letters []byte, target byte) byte {
    if target == 'z' {
        return letters[0]
    }
    for i := 0; i < len(letters); i++ {
        if letters[i] > target {
            return letters[i]
        }
    }
    return letters[0]
}