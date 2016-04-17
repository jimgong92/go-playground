// Echo4 prints out the CL arguments line by line alongside the index
package main

import (
  "fmt"
  "os"
)

func main() {
  for index, arg := range os.Args[1:] {
    fmt.Printf("%d. %s\n", index, arg)
  }
}