class Solution {
    fun isAnagram(s: String, t: String): Boolean {
        if (s.length != t.length) return false

        val array = IntArray(26) {0}
        for (i in s.indices){
            val posForS = s[i].code - 97
            val posForT = t[i].code - 97
            array[posForS]++
            array[posForT]--
        }
        
        for (i in array){
            if(i != 0)
                return false
        }
        return true
    }
}