class Solution {
    fun finalValueAfterOperations(operations: Array<String>): Int {
        var X = 0
    for (i in operations) {
        when (i) {
            "++X", "X++" -> X+=1
            "--X", "X--" -> X-=1
        }
    }
    return X
    }
}