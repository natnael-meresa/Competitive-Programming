from collections import Counter
k = int(input())
string = input()
 
letters = Counter(string)
ans = ["" for x in range(k)]
 
right = True
 
for x in letters:
    if letters[x]%k != 0:
        right = False
        break
    mul = letters[x]//k
    for i in range(k):
        ans[i] += x * mul
 
if not right:
    print(-1)
else:
    print("".join(ans))