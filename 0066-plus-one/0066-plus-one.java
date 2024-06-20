class Solution {
    public int[] plusOne(int[] digits) {
        int n = digits.length;
        
        //iterates the digits array from end to the start
        for (int i = n - 1; i >= 0; i--){
            //if thee digits less than 9, increment it
            if (digits[i] < 9) {
                digits[i]++;
                return digits;
            }
            //if digits is 9 then set to 0, because 9 + 0 = 10 which is [1, 0]
            digits[i] = 0;
        }
        //creates the new array with n+1 because the digits is 9 or 99 the array size will +1 [9,9] = [1,0,0] or [9] = [1,0]
        int[] newDigits = new int[n+1];
        //set the first element of array to 1
        newDigits[0] = 1;
        return newDigits;
    }
}