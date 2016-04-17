package main

import "fmt"

func main() {
  var a [5]int
  fmt.Println("emp:", a)

  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

  fmt.Println("len:", len(a))

  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)

  var twoD [2][3]int
  for i, v := range twoD {
    for j, _ := range v {
      twoD[i][j] = i * len(v) + j
    } 
  }
  fmt.Println("2d: ", twoD)
}