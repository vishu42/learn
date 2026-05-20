from stack import Stack


def is_balanced(input_str):
  stack = Stack()
  open, close = "(", ")"
  for item in input_str:
    if(item==open):
      print("---------------item==open--------------")
      print("stack size before pushing", stack.size())
      print("peeking", stack.peek())
      print("pushing: ", open)
      stack.push("whocareswhat")
      print("stack size after pushing", stack.size())
    elif(item==close):
      print("---------------item==close-------------")
      print("stack size before pop", stack.size())
      el = stack.pop()
      print("popped element", el)
      print("stack size after pop", stack.size())
      if(el == None):
        return False
      print("peeking after popping", stack.peek())
  if(stack.size()==0):
    print("stack size is zero")
    return True
  else:
    print("-------------------------------------")
    print("stack size in non-zero")
    print("stack size ",stack.size() )
    print("peeking", stack.peek())
    return False
