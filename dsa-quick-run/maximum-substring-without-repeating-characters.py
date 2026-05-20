def sol(str):
    seen = []
    best = 0
    for i in range(len(str)):
        while str[i] in seen:
            seen = seen[1:]
        seen.append(str[i])
        if len(seen)>best:
            best = len(seen)
    return best

print(sol("abcabcbb"))
print(sol("pwwkew"))