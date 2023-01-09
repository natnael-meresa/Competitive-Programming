t = int(input())
 
for i in range(t):
    n, c = input().split(" ")
    n = int(n)
   
    
    s = input()
    s = s+s
    
    bm = 0
    nxtg = 0
    for i in range(n):
        if s[i] == c:
            if i < nxtg:
                bm = max(bm, nxtg-i)
            else:
                for j in range(i+1,len(s)):
                    if s[j] == 'g':
                        nxtg = j
                        bm = max(bm, j-i)
                        break
    if c == 'g':
        print(0)
    else:
        print(bm)