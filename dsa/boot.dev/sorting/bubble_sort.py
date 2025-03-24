def bubble_sort(num):
  # [ 6, 5, 4, 2, 8, 9 ]
  if (len(num) == 0):
    return []
  if (len(num) == 1):
    return num
  swapping = True
  while(swapping):
    swapping = False
    for i in range(0, len(num)-1):
      print(i)
      if num[i] > num[i+1]:
        swap = num[i]
        num[i] = num[i+1]
        num[i+1] = swap
        swapping = True
  return num
