class Solution:
    def partitionLabels(self, s: str) -> List[int]:
       ans = []
       for x in s:
           isPart = False 
           for i in range(len(ans)):
               if x in ans[i]:
                   isPart = True
                   temp = ans[i:]
                   ans = ans[:i]
                   ans.append("".join(temp)+x)
                   break
           if not isPart:
               ans.append(x)
       finalans = [len(x) for x in ans] 
       return finalans
