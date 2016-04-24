package main

import "fmt"

type node struct {
  value interface{}
  next *node
}

type Stack struct {
  top *node
  size int
}

func (s *Stack) push(v interface{}) {
  tnode := node{value: v, next: s.top}
  s.top = &tnode
  s.size++
}

func (s *Stack) peek() interface{} {
  return s.top.value
}

func (s *Stack) pop() interface{} {
  tval := s.top.value
  s.top = s.top.next
  s.size--
  return tval
}

func main() {
  stack := Stack{}
  stack.push(1)
  stack.push(3)
  stack.push(5)
  fmt.Println(stack.peek() == 5)
  stack.push(7)
  fmt.Println(stack.pop() == 7)
  fmt.Println(stack.pop() == 5)
  fmt.Println(stack.pop() == 3)
  fmt.Println(stack.pop() == 1)


}