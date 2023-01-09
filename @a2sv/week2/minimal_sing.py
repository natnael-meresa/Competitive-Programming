n = int(input())

arr = []
for i in range(n):
    arr.append(input())

ans = []
for i in range(len(arr[0])):
    s = set()
    for x in range(n):
        c = arr[x][i]
        if c != "?":
            s.add(c)
    if len(s) > 1:
        ans.append("?")
    elif len(s) == 1:
        ans.append(s.pop())
    else:
        ans.append("x")
print("".join(ans))