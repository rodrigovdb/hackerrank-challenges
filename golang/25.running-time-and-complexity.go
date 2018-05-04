package main
import (
  "fmt"
  "math"
)

type number struct {
  value int
}

func read_from_input() []number{
  var size  int
  fmt.Scan(&size)

  items := make([]number, size)

  input := 0
  for i := 0; i < size; i++ {
    fmt.Scan(&input)
    items[i].value = input
  }

  return items
}

func (n number) show() {
  if n.value == 2 {
    fmt.Println("Prime")
    return
  }

  if n.value == 1 || n.value % 2 == 0 {
    fmt.Println("Not prime")
    return
  }

  for i := 3; i < int(math.Sqrt(float64(n.value))) + 1; i += 2 {
    if(n.value % i == 0){
      fmt.Println("Not prime")
      return
    }
  }

  fmt.Println("Prime")
}

func main() {
  for _, item := range read_from_input() {
    item.show()
  }
}
