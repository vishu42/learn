def sol(num, k):
    best = 0
    for i in range(len(num)-k + 1):
        total = 0
        window = []
        for j in range(k):
            total = total + num[i+j]
            window.append(num[i+j])
        if total > best:
            best = total
    return best


print(sol([2, 1, 5, 1, 3, 2], 3))