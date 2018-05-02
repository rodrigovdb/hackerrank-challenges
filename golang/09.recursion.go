package main
import "fmt"

func factorial(n int) int {
  if n > 1 {
    return n * factorial(n - 1)
  } else {
    return n
  }
}

func main() {
  var input int

  fmt.Scan(&input)

  fmt.Println(factorial(input))
}
