func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    mapC := map[byte]int{
        'a': 1,
        'b': 2,
        'c': 3,
        'd': 4,
        'e': 5,
        'f': 6,
        'g': 7,
        'h': 8,
    }
    
    if mapC[coordinate1[0]] % 2 == mapC[coordinate2[0]] % 2 {
        if (coordinate1[1] - '0') % 2 != (coordinate2[1] - '0') % 2 {
            return false
        }
    } else {
        if (coordinate1[1] - '0') % 2 == (coordinate2[1] - '0') % 2{
            return false
        }
    }
    return true 
}