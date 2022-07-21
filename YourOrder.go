func Order(sentence string) string {
  if sentence == "" {
    return ""
  }
  
  orderedArray := makeOrderedArray(sentence)
  solution := concateString(orderedArray)

  return solution
}

func makeOrderedArray(sentence string) []string {
  var orderredArray = make([]string, 9)
  var prevSpace int
  var stringPlace int
  
  for i, letter := range sentence {
    if letter > 48 && letter < 58 {
      stringPlace = int(letter) - 49
    }
    
    if i == len(sentence) - 1 {
      orderredArray[stringPlace] = sentence[prevSpace:i+1]
      break
    }

    if string(letter) == " " {
      orderredArray[stringPlace] = sentence[prevSpace:i]
      prevSpace = i + 1
      continue
    }
  }
  
  return orderredArray
}

func concateString(orderredArray []string) string {
  var orderedSentence string

  for _, value := range orderredArray {
    if value != "" {
      orderedSentence += value
      orderedSentence += " "
      continue
    }
    orderedSentence = orderedSentence[:len(orderedSentence) - 1]
    fmt.Println(len(orderedSentence))
    break
  }
  return orderedSentence
}
