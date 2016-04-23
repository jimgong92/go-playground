/**
 * Given a single input string, write a function that produces all possible anagrams
 * of a string and outputs them as an array
 * Ex.
 *  anagrams := allAnagrams("abc")
 *  fmt.Println(anagrams) // [ "abc" "acb" "bac" "bca" "cab" "cba" ]
 */
package main

import "fmt"

func main() {
  str := "abc"
  fmt.Println(allAnagrams(str))
}

func allAnagrams(str string) []string {
  var solutions = make(map[string]bool)
  return *permutation(str, solutions, "", &[]string{})
}

func permutation(options string, solutions map[string]bool, seq string, res *[]string) *[]string {
  if len(options) > 0 {
    for i, char := range options {
      updatedOptions := options[:i] + options[i+1:]
      updatedSeq := seq + string(char)
      permutation(updatedOptions, solutions, updatedSeq, res)
    }
  } else {
    if _, present := solutions[seq]; !present {
      solutions[seq] = true
      *res = append(*res, seq)
    }
  }
  return res
}