from collections import defaultdict
from collections import Counter

class Solution:
    def equalFrequency(self, word: str) -> bool:        
        letters = Counter(word)

        count = {}
        count = defaultdict(lambda:0,count)
        for x in letters.values():
            count[x] += 1

        if len(count) > 2:
            return False

        if len(count) == 1:
            for x in count:
                if x == 1 or count[x] == 1:
                    return True
        
        big = 0
        for x in count:
            if x > big:
                big  = x
                
        small = 0
        for x in count:
            if x < big:
                small = x

        if small == 1 and count[small] == 1:
            return True
        if big-1 == small and count[big] == 1:
            return True

        return False
            