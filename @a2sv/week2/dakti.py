t = int(input())

for i in range(t):
    s = input().split()
    mp = {}
    for x in range(len(s)):
        for u in range(len(s[x])):
            
            if s[x][u].isnumeric():
                il = u
                num = s[x][u]
                u += 1                
                while u < len(s[x]) and s[x][u].isnumeric():
                    num += s[x][u]
                    u += 1
                    
                
                mp[x] = int(num)
                s[x] = s[x][:il] + s[x][u:]
                break
    rst = [x[0] for x in sorted(mp.items(), key=lambda x:x[1])]
    finalStr = s[rst[0]]
    for i in range(1, len(rst)):
        finalStr += " " + s[rst[i]]
    print(finalStr)