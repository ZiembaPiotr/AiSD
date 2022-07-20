package kata

import (
  "strings"
)

func DuplicateEncode(word string) string {
  var solution string
  var lowerWord string
  numLetter := make(map[rune]int)
  
  lowerWord = strings.ToLower(word)
  
  for _, letter := range lowerWord {
    numLetter[letter] += 1
  }
  
  for _, letter := range lowerWord {
    if numLetter[letter] > 1 {
      solution += ")"
      continue
    }
    solution += "("
  }
  
  return solution 
}
