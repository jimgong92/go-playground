// Not actually covered by Go By Example, but by A Tour of Go
package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

type Person struct {
  name string
}

func main() {
  do(21)
  do("hello")
  do(true)
  do(Person{"Bob"})
}
