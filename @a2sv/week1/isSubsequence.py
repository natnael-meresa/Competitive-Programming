class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        tp = 0
        for x in s:
            found = False
            if tp == len(t):
                return False
            for i in range(tp, len(t)):
                if x == t[i]:
                    found = True
                    tp = i+1
                    break
            if not found:
                return False
        return True