/**
 * Reads a string from STDIN and print using Printf.
 */
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  reader  := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')

  fmt.Printf("Hello, World.\n%s", text)
}
