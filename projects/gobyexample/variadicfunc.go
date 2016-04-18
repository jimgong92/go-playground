package main

import "fmt"

func sum(nums ...int) (total int) {
  for _, num := range nums {
    total += num
  }
  return total
}

func main() {
  fmt.Println(sum(1, 2))
  fmt.Println(sum(1, 2, 3))
  nums := []int{1,2,3,4}
  fmt.Println(sum([]int{1,2,3,4, 5}...))
  fmt.Println(sum(nums...))
  fmt.Println(sum(append(nums, 10)...))
}