// Not actually covered by Go By Example, but by A Tour of Go
package main

import "fmt"

func main() {
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  // If code below is uncommented, will panic
  // f = i.(float64) // panic, unhandled assertion success
  // fmt.Println(f)
}
