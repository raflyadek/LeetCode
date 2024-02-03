class Solution {
    fun defangIPaddr(address: String): String {
        var result = ""
        for (i in address) if (i=='.') result+="[.]" else result += i
        return result
    }
}