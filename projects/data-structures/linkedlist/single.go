package main

import "fmt"

type node struct {
  value interface{}
  next *node
}

type SingleLL struct {
  head *node
  tail *node
}

func (l *SingleLL) addToTail(v interface{}) {
  tnode := &node{value: v}
  if (l.head == nil) {
    l.head = tnode
    l.tail = tnode
  } else {
    l.tail.next = tnode
    l.tail = l.tail.next
  }
}

func (l *SingleLL) removeHead() interface{} {
  if (l.head == nil) {
    return nil
  }
  hval := l.head.value
  l.head = l.head.next
  if (l.head == nil) {
    l.tail = nil
  }
  return hval
}

func (l *SingleLL) contains(v interface{}) bool {
  curr := l.head
  for curr != nil {
    if curr.value == v {
      return true
    }
    curr = curr.next;
  }
  return false
}

func main() {
  list := SingleLL{}
  list.addToTail(5)
  list.addToTail(6)
  list.addToTail(7)
  fmt.Println(list.removeHead())
  fmt.Println(list.removeHead())
  fmt.Println(list.removeHead())
}