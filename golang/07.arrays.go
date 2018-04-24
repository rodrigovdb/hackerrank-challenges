package main
import "fmt"

func readSize() int {
  var size int

    fmt.Scan(&size)

    return size
}

func readSlice(size int) []int {
  var elem int
  var arr []int

  for i := 0; i < size; i++ {
    fmt.Scan(&elem)
    arr = append(arr, elem)
  }

  return arr
}


func printSlice(arr []int) {
  for i := len(arr) - 1; i >= 0; i-- {
    fmt.Printf("%d ", arr[i])
  }
}

func main() {
	size  := readSize()
  arr   := readSlice(size)

  printSlice(arr)
}
