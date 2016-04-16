// Echo0 prints its CL arguments
package main

import (
  "fmt"
  "os"
)

func main() {
  s := os.Args[1]
  sep := ' '
  for _, arg := range os.Args[2:] {
    s += sep + arg
  }
  fmt.Println(s)
}