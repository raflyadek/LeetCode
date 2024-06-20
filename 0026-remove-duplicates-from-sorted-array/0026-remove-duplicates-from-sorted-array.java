class Solution {
    public int removeDuplicates(int[] nums) {
        // if nums length 0 return 0
        if (nums.length ==0) return 0;
        
        //create pointer i to mark the unique element in nums
        int i = 0;
        //pointer j will check the array from index 1 to the end of the array
        for(int j = 1; j < nums.length; j++){
            //check if there are duplicates
            if (nums[j] != nums[i]) {
                // if no duplicate index i will increment and update the nums[i]
                i++;
                nums[i] = nums[j];
            }
        }
        //after the loop finish the the new length of array without duplicates is i+1
        return i + 1;
    }
}