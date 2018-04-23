package main
import "fmt"

func display(number int) {
  for i := 1; i <= 10; i++ {
    fmt.Printf("%d x %d = %d\n", number, i, number * i);
  }
}

func main() {
  var input int

  fmt.Scan(&input)

  display(input)
}
