package main

import "fmt"

type node struct {
  value interface{}
  next *node
  prev *node
}

type DoubleLL struct {
  head *node
  tail *node
}

func (l *DoubleLL) addToTail(v interface{}) {
  tnode := &node{value: v, prev: l.tail}
  if (l.head == nil) {
    l.head = tnode
  } else {
    l.tail.next = tnode
  }
  l.tail = tnode
}

func (l *DoubleLL) removeFomTail() interface{} {
  if (l.tail == nil) {
    return nil
  }
  tval := l.tail.value
  l.tail = l.tail.prev
  if (l.tail == nil) {
    l.head = nil
  }
  return tval
}

func (l *DoubleLL) addToHead(v interface{}) {
  hnode := &node{value: v, next: l.head}
  if (l.head == nil) {
    l.tail= hnode
  } else {
    l.head.prev = hnode
  }
  l.head = hnode
}

func (l *DoubleLL) removeHead() interface{} {
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

func (l *DoubleLL) contains(v interface{}) bool {
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
  list := DoubleLL{}
  list.addToTail(5)
  list.addToTail(6)
  list.addToTail(7)
  list.addToHead(4)
  fmt.Println(list.removeHead())
  fmt.Println(list.removeHead())
  fmt.Println(list.removeHead())
}