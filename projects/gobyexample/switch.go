package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2
  fmt.Print("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  }

  t := time.Now()
  switch t.Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("it's the weekend")
  default:
    fmt.Println("it's a weekday")
  }

  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }
}