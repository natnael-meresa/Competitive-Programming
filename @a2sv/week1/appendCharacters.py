class Solution:
    def appendCharacters(self, s: str, t: str) -> int:
        tp = 0
        for x in range(len(t)):
            found = False
            if tp == len(t):
                return len(t)-x
            for i in range(tp, len(s)):
                if t[x] == s[i]:
                    found = True
                    tp = i+1
                    break
            if not found:
                return len(t)-x
        return 0