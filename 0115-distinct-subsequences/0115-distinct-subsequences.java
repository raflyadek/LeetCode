class Solution {
    public int numDistinct(String s, String t) {
        //length for target string 't'
        int lengthTarget = t.length();
        
        //dp array for sorting the number
        int[] dp = new int[lengthTarget + 1];
        
        //base case initialization
        dp[0] = 1;
        
        //iterate over the string s
        for (char sourceChar : s.toCharArray()) {
            //iterate backward through the dp array 
            for(int j = lengthTarget; j > 0; --j) {
                //get the jth character of string target t
                char targetChar = t.charAt(j-1);
                //if the current character in 's' and 't' matches
                //we add the number to the previous character
                if (sourceChar == targetChar) {
                    dp[j] += dp[j-1];
                }
            }
        }
        return dp[lengthTarget];
    }
}