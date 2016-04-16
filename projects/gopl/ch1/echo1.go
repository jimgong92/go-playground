// Echo1 prints its command-line arguments
package main

import (
  "fmt"
  "os"
)

func main() {
  var s, sep string // Uses default values of empty strings
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}
