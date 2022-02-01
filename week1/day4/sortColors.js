/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var sortColors = function(nums) {
    for (let i = 1; i < nums.length; i++) {
            // Choosing the first element in our unsorted subarray
            let current = nums[i];
            // The last element of our sorted subarray
            let j = i-1; 
            while ((j > -1) && (current < nums[j])) {
                nums[j+1] = nums[j];
                j--;
            }
            nums[j+1] = current;
        }
};