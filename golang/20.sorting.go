package main
import "fmt"

type data struct {
  swaps   int
  items   []int
}

func new_data() data {
  var size int
  fmt.Scan(&size)

  item := data { items: make([]int, size) }

  for i := 0; i < size; i++ {
    fmt.Scan(&item.items[i])
  }

  return item
}

func (d data) sort() data {
  d.swaps = 0
  aux  := 0

  for {
    changes := false

    for i := 0; i < len(d.items) - 1; i++ {
      if d.items[i] > d.items[i + 1]{
        aux       = d.items[i]
        d.items[i]    = d.items[i + 1]
        d.items[i + 1]  = aux

        changes = true
        d.swaps += 1
      }
    }

    if !changes {
      break
    }
  }

  return d
}

func (d data) first_element() int {
  return d.items[0]
}

func (d data) last_element() int {
  return d.items[len(d.items) - 1]
}

func (d data) show() {
  fmt.Printf("Array is sorted in %d swaps.\n", d.swaps)
  fmt.Printf("First Element: %d\n", d.first_element())
  fmt.Printf("Last Element: %d", d.last_element())
}

func main() {
  new_data().sort().show()
}
