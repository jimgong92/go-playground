package main

import (
  "fmt"
  "time"
)

func main() {
  c1 := make(chan string)
  go func() {
    time.Sleep(time.Second * 2)
    c1 <- "Result 1"
  }()

  select {
  case res := <-c1:
    fmt.Println(res)
  case <-time.After(time.Second):
    fmt.Println("timeout 1")
  }

  c2 := make(chan string)
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "Result 2"
  }()

  select {
  case res := <-c2:
    fmt.Println(res)
  case <-time.After(time.Second * 3):
    fmt.Println("timeout 2")
  }
}