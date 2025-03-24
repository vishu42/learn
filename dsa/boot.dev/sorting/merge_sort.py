def merge_sort(list):
  print("==================================")
  print("list->", list)
  if (len(list) < 2):
    return list
  mid = len(list)//2
  print("left list", list[0:mid])
  print("right list", list[mid: len(list)])
  left = merge_sort(list[0:mid])
  right = merge_sort(list[mid:len(list)])
  return merge(left, right)


def merge(a, b):
  res = []
  i,j = 0,0
  while(i < len(a) and j < len(b)):
    if (a[i] < b[j]):
      res.append(a[i])
      i = i + 1
    else:
      res.append(b[j])
      j = j + 1

  while (i < len(a)):
    res.append(a[i])
    i = i + 1
  while (j < len(b)):
    res.append(b[j])
    j = j + 1
  return res
