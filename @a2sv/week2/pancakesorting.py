class Solution:
    def pancakeSort(self, arr: List[int]) -> List[int]:
        sortd = arr.copy()
        sortd.sort()
        ans = []
        for i in range(len(arr)-1, -1 ,-1):
            for j in range(i, -1 ,-1):
                if sortd[i] == arr[j] and not i == j:
                    arr = arr[:j+1][::-1] + arr[j+1:]
                    arr = arr[:i+1][::-1] + arr[i+1:]
                    ans.append(j+1)
                    ans.append(i+1)
                    break
        return ans
