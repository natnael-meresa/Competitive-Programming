class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        nums.sort()
        ansNums = [0,0]
        nums = [0]+nums
        miset = False
        d = 0 
        for idx in range(1, len(nums)):
            if nums[idx] == nums[idx-1]:
                ansNums[0] = nums[idx]
                d += 1
            if idx-d != nums[idx] and not miset:
                miset = True
                ansNums[1] = idx-d
        if not miset:
            ansNums[1] = len(nums)-1
        return ansNums