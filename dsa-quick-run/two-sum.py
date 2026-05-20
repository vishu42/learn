def two_sum(numbers, target):
    # value -> index
    seen = {}
    for i, v in enumerate(numbers):
        needed = target - v
        if needed in seen:
            return [seen[needed], i]
        seen[v] = i
    return []


print(two_sum([3, 2, 4], 7))