package main
import "fmt"

func parse(word string) {
  var odds  string = ""
  var evens   string = ""

  for i := 0; i < len(word); i++ {
    if i % 2 == 0 {
      odds = odds + string(word[i])
    } else {
      evens = evens + string(word[i])
    }
  }

  fmt.Printf("%s %s", odds, evens);
}

func main() {
  var number int
  var word   string

  fmt.Scan(&number)

  for i := 0; i < number; i++ {
    fmt.Scan(&word)

    parse(word)
    fmt.Println()
  }
}
