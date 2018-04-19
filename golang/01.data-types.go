package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
    var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)

		/**
     * MY CODE STARTS HERE
     */
    // Declare second integer, double, and String variables.
    var my_integer uint64
    var my_double  float64
    var my_string  string

    // Read and save an integer, double, and String to your variables.
    fmt.Scan(&my_integer)
    fmt.Scan(&my_double)
    for scanner.Scan(){ my_string = scanner.Text() }

    // Print the sum of both integer variables on a new line.
    fmt.Printf("%d\n", i + my_integer)

    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f\n", d + my_double)

    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Printf("%s", s + my_string)
}
