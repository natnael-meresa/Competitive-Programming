class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        finalArr = [0] * len(nums)
        
        for i in range(len(nums)):
            for j in range(len(nums)):
                if(j != i and nums[j] < nums[i]):
                    finalArr[i] += 1
        return finalArr