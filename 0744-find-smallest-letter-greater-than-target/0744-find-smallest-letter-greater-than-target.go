func nextGreatestLetter(letters []byte, target byte) byte {
    if target == byte('z') {
        return letters[0]
    }
    for i := 0; i < len(letters); i++ {
        if target < letters[i] {
            return letters[i]
        }
    }
    return letters[0]
}