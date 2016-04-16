// Echo3 prints its CL arguments
package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}