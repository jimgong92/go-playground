package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type Rect struct {
  width, height int
}

func (r *Rect) String() string {
  return fmt.Sprintf("[width: %v, height: %v]", r.width, r.height)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a, z)
  
  r := &Rect{20, 10}
  fmt.Println(r)
}
