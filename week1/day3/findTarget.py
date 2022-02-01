class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()
        
        if target not in nums:
            return []
        indexs = []
        for i in range(len(nums)):
            if nums[i] == target:
                indexs.append(i)
        
        return indexs