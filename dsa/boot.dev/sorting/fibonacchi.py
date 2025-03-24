def fib(n):
  if(n==0):
    return 0
  if(n==1):
    return 1
  grandparent, parent, current = 0, 1, 0
  for i in range(0, n-1):
    current = parent + grandparent
    swap = parent
    parent = current
    grandparent = swap
  return current

print(fib(5))
