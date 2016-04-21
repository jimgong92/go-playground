/**
 * Given a single input string, write a function that produces all possible anagrams
 * of a string and outputs them as an array
 * Ex.
 *  anagrams := allAnagrams("abc")
 *  fmt.Println(anagrams) // [ "abc" "acb" "bac" "bca" "cab" "cba" ]
 */
package main

import (
  "fmt"
  "strings"
)

func main() {
  str := "abc"
  fmt.Println(allAnagrams(str))
}

func allAnagrams(str string) []string {
  var anagrams []string
  solutions := make(map[string]bool)
  options := strings.Split(str, "")
  permutation([]string{}, options, &anagrams, solutions)
  return anagrams
}

func permutation(seq []string, options []string, res *[]string, solutions map[string]bool) {
  if len(options) > 0 {
    for i, char := range options {
      duplOpt := append([]string{}, options...)
      duplOpt = append(duplOpt[:i], duplOpt[i+1:]...)
      duplSeq := append([]string{}, seq...)
      duplSeq = append(duplSeq, char)
      permutation(duplSeq, duplOpt, res, solutions)
    }
  } else {
    solution := strings.Join(seq, "")
    if _, present := solutions[solution]; !present {
      solutions[solution] = true
      *res = append(*res, solution)
    }
  }
}