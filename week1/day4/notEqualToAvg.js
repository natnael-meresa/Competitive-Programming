/**
 * @param {number[]} nums
 * @return {number[]}
 */
var rearrangeArray = function(nums) {
    
    nums.sort((a,b) => a -b)
    let left = 0;
    let right = nums.length -1
    let final = []
    while(final.length != nums.length){
        final.push(nums[left]);
        left++
        if (left <= right){
            final.push(nums[right]);
            right--;
        }
    }
    
    
    return final
};